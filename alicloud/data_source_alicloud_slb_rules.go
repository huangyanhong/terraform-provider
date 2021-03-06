package alicloud

import (
	"fmt"
	"log"
	"regexp"

	"github.com/alibaba/terraform-provider/alicloud/connectivity"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAlicloudSlbRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlicloudSlbRulesRead,

		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"frontend_port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
				MinItems: 1,
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateNameRegex,
				ForceNew:     true,
			},

			// Computed values
			"slb_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAlicloudSlbRulesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	args := slb.CreateDescribeRulesRequest()
	args.LoadBalancerId = d.Get("load_balancer_id").(string)
	args.ListenerPort = requests.NewInteger(d.Get("frontend_port").(int))

	idsMap := make(map[string]string)
	if v, ok := d.GetOk("ids"); ok {
		for _, vv := range v.([]interface{}) {
			idsMap[Trim(vv.(string))] = Trim(vv.(string))
		}
	}

	raw, err := client.WithSlbClient(func(slbClient *slb.Client) (interface{}, error) {
		return slbClient.DescribeRules(args)
	})
	if err != nil {
		return fmt.Errorf("DescribeRules got an error: %#v", err)
	}
	resp, _ := raw.(*slb.DescribeRulesResponse)
	if resp == nil {
		return fmt.Errorf("there is no SLB with the ID %s. Please change your search criteria and try again", args.LoadBalancerId)
	}

	var filteredRulesTemp []slb.Rule
	nameRegex, ok := d.GetOk("name_regex")
	if (ok && nameRegex.(string) != "") || (len(idsMap) > 0) {
		var r *regexp.Regexp
		if nameRegex != "" {
			r = regexp.MustCompile(nameRegex.(string))
		}
		for _, rule := range resp.Rules.Rule {
			if r != nil && !r.MatchString(rule.RuleName) {
				continue
			}
			if len(idsMap) > 0 {
				if _, ok := idsMap[rule.RuleId]; !ok {
					continue
				}
			}

			filteredRulesTemp = append(filteredRulesTemp, rule)
		}
	} else {
		filteredRulesTemp = resp.Rules.Rule
	}

	if len(filteredRulesTemp) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	log.Printf("[DEBUG] alicloud_slb_rules - Slb rules found: %#v", filteredRulesTemp)

	return slbRulesDescriptionAttributes(d, filteredRulesTemp)
}

func slbRulesDescriptionAttributes(d *schema.ResourceData, rules []slb.Rule) error {
	var ids []string
	var s []map[string]interface{}

	for _, rule := range rules {
		mapping := map[string]interface{}{
			"id":              rule.RuleId,
			"name":            rule.RuleName,
			"domain":          rule.Domain,
			"url":             rule.Url,
			"server_group_id": rule.VServerGroupId,
		}

		log.Printf("[DEBUG] alicloud_slb_rules - adding slb_rule mapping: %v", mapping)
		ids = append(ids, rule.RuleId)
		s = append(s, mapping)
	}

	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("slb_rules", s); err != nil {
		return err
	}

	// create a json file in current directory and write data source to it.
	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		writeToFile(output.(string), s)
	}
	return nil
}

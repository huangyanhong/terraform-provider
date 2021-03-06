package teslamaxcompute

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// SaleInfo is a nested struct in teslamaxcompute response
type SaleInfo struct {
	SaleMode    string `json:"SaleMode" xml:"SaleMode"`
	Uid         string `json:"Uid" xml:"Uid"`
	Mem         int    `json:"Mem" xml:"Mem"`
	Cpu         int    `json:"Cpu" xml:"Cpu"`
	BizCategory string `json:"BizCategory" xml:"BizCategory"`
	Owner       string `json:"Owner" xml:"Owner"`
	QueryDate   string `json:"QueryDate" xml:"QueryDate"`
}

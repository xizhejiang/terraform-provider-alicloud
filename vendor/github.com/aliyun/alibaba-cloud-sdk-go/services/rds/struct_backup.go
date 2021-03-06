package rds

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

// Backup is a nested struct in rds response
type Backup struct {
	BackupId                  string `json:"BackupId" xml:"BackupId"`
	DBInstanceId              string `json:"DBInstanceId" xml:"DBInstanceId"`
	BackupStatus              string `json:"BackupStatus" xml:"BackupStatus"`
	BackupStartTime           string `json:"BackupStartTime" xml:"BackupStartTime"`
	BackupEndTime             string `json:"BackupEndTime" xml:"BackupEndTime"`
	BackupType                string `json:"BackupType" xml:"BackupType"`
	BackupMode                string `json:"BackupMode" xml:"BackupMode"`
	BackupMethod              string `json:"BackupMethod" xml:"BackupMethod"`
	BackupDownloadURL         string `json:"BackupDownloadURL" xml:"BackupDownloadURL"`
	BackupIntranetDownloadURL string `json:"BackupIntranetDownloadURL" xml:"BackupIntranetDownloadURL"`
	BackupLocation            string `json:"BackupLocation" xml:"BackupLocation"`
	BackupExtractionStatus    string `json:"BackupExtractionStatus" xml:"BackupExtractionStatus"`
	BackupScale               string `json:"BackupScale" xml:"BackupScale"`
	BackupDBNames             string `json:"BackupDBNames" xml:"BackupDBNames"`
	TotalBackupSize           int64  `json:"TotalBackupSize" xml:"TotalBackupSize"`
	BackupSize                int64  `json:"BackupSize" xml:"BackupSize"`
	HostInstanceID            string `json:"HostInstanceID" xml:"HostInstanceID"`
	StoreStatus               string `json:"StoreStatus" xml:"StoreStatus"`
	MetaStatus                string `json:"MetaStatus" xml:"MetaStatus"`
	SlaveStatus               string `json:"SlaveStatus" xml:"SlaveStatus"`
	ConsistentTime            int64  `json:"ConsistentTime" xml:"ConsistentTime"`
}

/*
Copyright © 2020 PT Bhinneka Mentari Dimensi.

NOTICE:

All Information contained herein is, and remains the property of PT Bhinneka Mentari Dimensi and its suppliers, if any.
The intellectual and technical concepts contained herein are proprietary to PT Bhinneka Mentari Dimensi and its suppliers
and maybe covered by Republic of Indonesia and Foreign Patents, patents in process, and are protected by trade secret or copyright law.
Dissemination of this information or reproduction of this material is strictly forbidden
unless prior written permission is obtained from PT Bhinneka Mentari Dimensi
*/

// Package record of shared (version 1) is a package that handle anything about data object record such as :
//
// * embedded status record model
// 	- userIn which represent id of the user who insert the record
// 	- userUp which represent id of the user who update the record
// 	- dateIn which represent datetime of when the record inserted for the first time
// 	- dateUp which represent datetime of when is the last time record was updated / deleted
// 	- statusRecord which represent status of the record ("N" = new record, "U" = updated record, "D" = deleted record)
//
// * record validator to validate value from database record and return it as nil / null if the value is empty, zero time, or zero integer
package record

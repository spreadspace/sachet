/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type SetChatStatusInputObject struct {
	// Chat ID.
	Id int32 `json:"id,omitempty"`
	// Chat status:   * **a** - Active;   * **c** - Closed;   * **d** - Deleted. 
	Status string `json:"status,omitempty"`
}

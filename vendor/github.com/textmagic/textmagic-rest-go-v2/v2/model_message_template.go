/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type MessageTemplate struct {
	// Template ID.
	Id int32 `json:"id"`
	// Template name.
	Name string `json:"name"`
	// Template text. May contain tags inside braces. See the [Custom fields list](https://docs.textmagic.com/#section/Custom-fields-list-(Merge-tags)).
	Content string `json:"content"`
	// Time when the template was last modified.
	LastModified string `json:"lastModified"`
}

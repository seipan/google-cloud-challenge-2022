// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/estate"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/etype"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/event"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/schema"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	estateFields := schema.EState{}.Fields()
	_ = estateFields
	// estateDescName is the schema descriptor for name field.
	estateDescName := estateFields[0].Descriptor()
	// estate.NameValidator is a validator for the "name" field. It is called by the builders before save.
	estate.NameValidator = estateDescName.Validators[0].(func(string) error)
	etypeFields := schema.EType{}.Fields()
	_ = etypeFields
	// etypeDescName is the schema descriptor for name field.
	etypeDescName := etypeFields[0].Descriptor()
	// etype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	etype.NameValidator = etypeDescName.Validators[0].(func(string) error)
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescName is the schema descriptor for name field.
	eventDescName := eventFields[0].Descriptor()
	// event.DefaultName holds the default value on creation for the name field.
	event.DefaultName = eventDescName.Default.(string)
	// event.NameValidator is a validator for the "name" field. It is called by the builders before save.
	event.NameValidator = func() func(string) error {
		validators := eventDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// eventDescDetail is the schema descriptor for detail field.
	eventDescDetail := eventFields[1].Descriptor()
	// event.DetailValidator is a validator for the "detail" field. It is called by the builders before save.
	event.DetailValidator = eventDescDetail.Validators[0].(func(string) error)
	// eventDescLocation is the schema descriptor for location field.
	eventDescLocation := eventFields[2].Descriptor()
	// event.LocationValidator is a validator for the "location" field. It is called by the builders before save.
	event.LocationValidator = eventDescLocation.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[0].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = func() func(int) error {
		validators := userDescAge.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(age int) error {
			for _, fn := range fns {
				if err := fn(age); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescAuthenticated is the schema descriptor for authenticated field.
	userDescAuthenticated := userFields[2].Descriptor()
	// user.DefaultAuthenticated holds the default value on creation for the authenticated field.
	user.DefaultAuthenticated = userDescAuthenticated.Default.(bool)
	// userDescGmail is the schema descriptor for gmail field.
	userDescGmail := userFields[3].Descriptor()
	// user.GmailValidator is a validator for the "gmail" field. It is called by the builders before save.
	user.GmailValidator = userDescGmail.Validators[0].(func(string) error)
	// userDescIconImg is the schema descriptor for icon_img field.
	userDescIconImg := userFields[4].Descriptor()
	// user.IconImgValidator is a validator for the "icon_img" field. It is called by the builders before save.
	user.IconImgValidator = userDescIconImg.Validators[0].(func(string) error)
}

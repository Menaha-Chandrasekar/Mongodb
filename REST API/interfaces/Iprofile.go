/*
5 method Signatures
Creating the profile
EditOne profile
EditMany Profile
deleting profile
Fetch profile
this interface will provide the required methods
*/
package interfaces 

import "restapi/models"



type Iprofile interface{
        CreateProfile(profile *models.Profile)
		Editprofile(profile *models.Profile)
		DeleteProfile(profileId int)
		GetProfileById(profile int)
		GetProfileBySearch()
}
/* It contains the implementation from the interfaces*/

package services

import (
	"context"
	"restapi/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileService struct{
	ProfileCollection *mongo.Collection
	ctx context.Context
}
func (p *ProfileService) CreateProfile(profile *models.Profile){}
func (p *ProfileService) EditProfile(profile *models.Profile){}
func (p *ProfileService) DeleteProfile(profileId int){}
func (p *ProfileService) GetProfileById(profileId int){}
func (p *ProfileService) GetProfileBySearch(){}
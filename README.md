# travel_form_approval
This is a backend API for Travel form approval using Golang.
In this system, users can submit forms to request travel plans for work.

Their superiors will be able to approve or reject these plans, and the approval flow strecthes for a few approvals. 

This system uses PostgreSQL as the database, and GORM for the ORM.

Currently implemented:
- Test endpoint
- Login with JWT
- Create user with validations (credential and JWT)
- Get Users with JWT validation

For the API documentatiom, please refer to this link: https://www.postman.com/tastycarbonara/workspace/tastycarbonara/collection/33251652-174a8713-362a-4992-be35-3815ade43043?action=share&creator=33251652
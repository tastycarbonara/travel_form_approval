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

For the API documentatiom, please refer to this link: https://smaga-lms.postman.co/workspace/My-Workspace~4bde0820-d16f-4f43-8529-6401c466743d/collection/36770127-b46a1bda-eee6-46d1-bf11-32b4969c5c7e?action=share&source=copy-link&creator=36770127
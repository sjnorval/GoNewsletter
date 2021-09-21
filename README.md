# GoNewsletter
### TodoList
1. Create Project (Done)
    * - [x] [Server](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/server.go)
2. Create Models
    * - [x] [User](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/models/User.go)
    * - [x] [Topic](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/models/Topic.go)
    * - [x] [NewsPost](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/models/NewsPost.go)
    * - [x] [UserTopic](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/models/User_Topic_Hambgr.go)
3. Model Utility 
    * - [x] [User](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/models/User.go)
    * - [x] [Topic](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/models/Topic.go) 
    * - [ ] [NewsPost](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/models/NewsPost.go)
      * Save() Required
      * Prepare() Required
4. Database Creation Scripts (Pending)
    * - [ ] Users
    * - [ ] Topics
    * - [ ] NewsPosts
    * - [ ] TopicsInUsers
5. MiddleWare Utilities
    * - [x] [ErrorFormatting](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/Utils/formaterror.go)
    * - [x] [Guid](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/Utils/guid.go)
      > Would have liked to use uuid alhough I have not dug into how to use it correctly within Go.
6. Controllers
    * - [x] [Base](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/controllers/base.go)
    * - [x] [Users_Controller](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/controllers/users_controller.go)
    * - [ ] [Topics_Controller](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/controllers/topic_controller.go)
    > [RegisterUserToTopic](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/controllers/topic_controller.go#L70) is not completed and still needs adjusting.
    * - [ ] [NewsPosts_Controller](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/controllers/newsPost_controller.go)
    > All functions need adjusting. Copied from a sample file just to not work on a blank canvas, I do it sometimes to make sure I don't overthink the problem.
    * - [x] [Routes](https://github.com/sjnorval/GoNewsletter/blob/6dfc4815c2d15a9fd9d84e750fd0759134b188e9/api/controllers/routes.go)

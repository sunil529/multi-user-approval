created a dummy web application for this whole scenario
Postgres Database and Login:
 2.
 ● Create a PostgresSQL database to store user information such as name
 and email.
 ● Develop a login APIs that allows users to sign up and assigns them an
 internal login ID.
 ● Develop APIs to create tasks. Any user can create a task. (You can
 assume a schema for the task.). Taks should have status.
 ● Implement a bonus feature That lets other users provide approval for
 the task
 2. Multi approval Process:
 ● Develop APIs for multi-approval processes where a task needs to be
 approved by 3 other users to be in approved status.
 ● Provide design pointers or hooks for
 ○ User creating task to choose the 3 other users from a dropdown
 list and send email notifications to each user when a new task is
 created.
 ○ Add the functionality for users to add comments.
 ○ Ensure the process creator receives a notification on their page
 when anyone signs off, and notify all parties involved via email
 when everyone signs off.
 3. API:
 ○ Break down the multi-signature process into REST APIs.
 ○ Ensure that the APIs can be integrated into any webpage

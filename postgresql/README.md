## Auto Message Sender Service Database

PostgreSQL database used for storing informations about messages and storing config for system behavior.

## init.sql

Upon installation, the creation of schemas, users, and tables is handled here, along with the insertion of rows into these tables.

### Tables

#### Config Table: 

| ID             | Name          | Status        |
| :------------- | :------------ | :------------ |
| **PK**, `UUID` | `VARCHAR(30)` | `VARCHAR(3)` |


#### Messages Table: 

| ID             | Content        | Recipient Phone Number | SentStatus    | MessageId | CreationTime | SentTime    |
| :------------- | :------------- | :--------------------- | ------------- | --------- | ------------ | ----------- |
| **PK**, `UUID` | `VARCHAR(300)` | `VARCHAR(20)`         | `VARCHAR(10)` | `UUID`    | `TIMESTAMP`  | `TIMESTAMP` |




  
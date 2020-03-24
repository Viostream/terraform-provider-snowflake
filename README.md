# Terraform Provider: Snowflake

----

[![Build Status](https://travis-ci.com/viostream/terraform-provider-snowflake.svg?branch=master)](
https://travis-ci.com/viostream/terraform-provider-snowflake)
[![codecov](https://codecov.io/gh/viostream/terraform-provider-snowflake/branch/master/graph/badge.svg)](
https://codecov.io/gh/viostream/terraform-provider-snowflake)

This is a terraform provider plugin for managing
[Snowflake](http://snowflakedb.com) accounts. Initially developed by
[Chan Zuckerberg](https://github.com/chanzuckerberg/terraform-provider-snowflake),
currently maintained by [Viostream](https://viostream.com).

## Install

```shell
go get -u github.com/viostream/terraform-provider-snowflake
# *NIX
mv $GOPATH/bin/terraform-provider-snowflake ~/.terraform.d/plugins
# Windo$e
???
```

## Authentication

<<<<<<< HEAD
<<<<<<< HEAD
We currently only support username + password auth and suggest that you only do
so via environment variables. So a config something like:
=======
We currently support username + password and keypair auth and suggest that you do so via environment variables. Define a config with something like-
>>>>>>> 7dd55fa02ff8b69235d11375c3fb5f2028e5146b
=======
We currently support username + password, browser and keypair authenthication. We suggest that you do so via environment variables. Define a config with the non-senstive field like-
>>>>>>> 4346dee4430764fa7c608a761b5ffd3c53650631

```hcl
provider "snowflake" {
  account = "..."
  role    = "..."
  region  = "..."
}
```

Then set `SNOWFLAKE_USER` and either `SNOWFLAKE_PASSWORD` or `SNOWFLAKE_PRIVATE_KEY_PATH`.

### Keypair Authentication Environment Variables
You should generate the public and private keys and set up environment variables.

```shell
cd ~/.ssh
openssl genrsa -out snowflake_key 4096
openssl rsa -in snowflake_key -pubout -out snowflake_key.pub
```

To export the variables into your provider:
```shell
export SNOWFLAKE_USER="..."
export SNOWFLAKE_PRIVATE_KEY_PATH="~/.ssh/snowflake_key"
```

### Username and Password Environment Variables
If you choose to use Username and Password Authentication, export these credentials:
```shell
export SNOWFLAKE_USER='...'
export SNOWFLAKE_PASSWORD='...'
```

## Resources

<<<<<<< HEAD
We support managing a subset of snowflakedb resources, with a focus on access
control and management.
=======
We support managing a subset of snowflakedb resources, with a focus on access control and management. We've built and support the resources we use. If you are lookig for others to be supported we are more than happy to get PRs merged.
>>>>>>> 4346dee4430764fa7c608a761b5ffd3c53650631

You can see a number of examples [here](examples).

<!-- START -->

### snowflake_database

#### properties

|            NAME             |  TYPE  |                                  DESCRIPTION                                  | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|-----------------------------|--------|-------------------------------------------------------------------------------|----------|-----------|----------|---------|
| comment                     | string |                                                                               | true     | false     | false    | ""      |
| data_retention_time_in_days | int    |                                                                               | true     | false     | true     | <nil>   |
| from_database               | string | Specify a database to create a clone from.                                    | true     | false     | false    | <nil>   |
| from_share                  | map    | Specify a provider and a share in this map to create a database from a share. | true     | false     | false    | <nil>   |
| name                        | string |                                                                               | false    | true      | false    | <nil>   |

### snowflake_database_grant

**Note**: The snowflake_database_grant resource creates exclusive attachments of grants.
Across the entire Snowflake account, all of the databases to which a single grant is attached must be declared
by a single snowflake_database_grant resource. This means that even any snowflake_database that have the attached
grant via any other mechanism (including other Terraform resources) will have that attached grant revoked by this resource.
These resources do not enforce exclusive attachment of a grant, it is the user's responsibility to enforce this.

#### properties

|     NAME      |  TYPE  |                      DESCRIPTION                       | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|---------------|--------|--------------------------------------------------------|----------|-----------|----------|---------|
| database_name | string | The name of the database on which to grant privileges. | false    | true      | false    | <nil>   |
| privilege     | string | The privilege to grant on the database.                | true     | false     | false    | "USAGE" |
| roles         | set    | Grants privilege to these roles.                       | true     | false     | false    | <nil>   |
| shares        | set    | Grants privilege to these shares.                      | true     | false     | false    | <nil>   |

### snowflake_managed_account

#### properties

|      NAME      |  TYPE  |                                                                  DESCRIPTION                                                                   | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT  |
|----------------|--------|------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|----------|
| admin_name     | string | Identifier, as well as login name, for the initial user in the managed account. This user serves as the account administrator for the account. | false    | true      | false    | <nil>    |
| admin_password | string | Password for the initial user in the managed account.                                                                                          | false    | true      | false    | <nil>    |
| cloud          | string | Cloud in which the managed account is located.                                                                                                 | false    | false     | true     | <nil>    |
| comment        | string | Specifies a comment for the managed account.                                                                                                   | true     | false     | false    | <nil>    |
| created_on     | string | Date and time when the managed account was created.                                                                                            | false    | false     | true     | <nil>    |
| locator        | string | Display name of the managed account.                                                                                                           | false    | false     | true     | <nil>    |
| name           | string | Identifier for the managed account; must be unique for your account.                                                                           | false    | true      | false    | <nil>    |
| region         | string | Snowflake Region in which the managed account is located.                                                                                      | false    | false     | true     | <nil>    |
| type           | string | Specifies the type of managed account.                                                                                                         | true     | false     | false    | "READER" |
| url            | string | URL for accessing the managed account, particularly through the web interface.                                                                 | false    | false     | true     | <nil>    |

### snowflake_pipe

#### properties

|         NAME         |  TYPE  |                                                   DESCRIPTION                                                   | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|----------------------|--------|-----------------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| auto_ingest          | bool   | Specifies a auto_ingest param for the pipe.                                                                     | true     | false     | false    | false   |
| comment              | string | Specifies a comment for the pipe.                                                                               | true     | false     | false    | <nil>   |
| copy_statement       | string | Specifies the copy statement for the pipe.                                                                      | false    | true      | false    | <nil>   |
| database             | string | The database in which to create the pipe.                                                                       | false    | true      | false    | <nil>   |
| name                 | string | Specifies the identifier for the pipe; must be unique for the database and schema in which the pipe is created. | false    | true      | false    | <nil>   |
| notification_channel | string | Amazon Resource Name of the Amazon SQS queue for the stage named in the DEFINITION column.                      | false    | false     | true     | <nil>   |
| owner                | string | Name of the role that owns the pipe.                                                                            | false    | false     | true     | <nil>   |
| schema               | string | The schema in which to create the pipe.                                                                         | false    | true      | false    | <nil>   |

### snowflake_resource_monitor

#### properties

|            NAME            |  TYPE  |                                                                   DESCRIPTION                                                                   | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|----------------------------|--------|-------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| credit_quota               | float  | The amount of credits allocated monthly to the resource monitor, round up to 2 decimal places.                                                  | true     | false     | true     | <nil>   |
| end_timestamp              | string | The date and time when the resource monitor suspends the assigned warehouses.                                                                   | true     | false     | false    | <nil>   |
| frequency                  | string | The frequency interval at which the credit usage resets to 0. If you set a frequency for a resource monitor, you must also set START_TIMESTAMP. | true     | false     | true     | <nil>   |
| name                       | string | Identifier for the resource monitor; must be unique for your account.                                                                           | false    | true      | false    | <nil>   |
| notify_triggers            | set    | A list of percentage thresholds at which to send an alert to subscribed users.                                                                  | true     | false     | false    | <nil>   |
| start_timestamp            | string | The date and time when the resource monitor starts monitoring credit usage for the assigned warehouses.                                         | true     | false     | true     | <nil>   |
| suspend_immediate_triggers | set    | A list of percentage thresholds at which to immediately suspend all warehouses.                                                                 | true     | false     | false    | <nil>   |
| suspend_triggers           | set    | A list of percentage thresholds at which to suspend all warehouses.                                                                             | true     | false     | false    | <nil>   |

### snowflake_role

#### properties

|  NAME   |  TYPE  | DESCRIPTION | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|---------|--------|-------------|----------|-----------|----------|---------|
| comment | string |             | true     | false     | false    | <nil>   |
| name    | string |             | false    | true      | false    | <nil>   |

### snowflake_role_grants

#### properties

|   NAME    |  TYPE  |              DESCRIPTION              | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|-----------|--------|---------------------------------------|----------|-----------|----------|---------|
| role_name | string | The name of the role we are granting. | false    | true      | false    | <nil>   |
| roles     | set    | Grants role to this specified role.   | true     | false     | false    | <nil>   |
| users     | set    | Grants role to this specified user.   | true     | false     | false    | <nil>   |

### snowflake_schema

#### properties

|        NAME         |  TYPE  |                                                                                                                      DESCRIPTION                                                                                                                       | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|---------------------|--------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| comment             | string | Specifies a comment for the schema.                                                                                                                                                                                                                    | true     | false     | false    | <nil>   |
| data_retention_days | int    | Specifies the number of days for which Time Travel actions (CLONE and UNDROP) can be performed on the schema, as well as specifying the default Time Travel retention time for all tables created in the schema.                                       | true     | false     | false    |       1 |
| database            | string | The database in which to create the schema.                                                                                                                                                                                                            | false    | true      | false    | <nil>   |
| is_managed          | bool   | Specifies a managed schema. Managed access schemas centralize privilege management with the schema owner.                                                                                                                                              | true     | false     | false    | false   |
| is_transient        | bool   | Specifies a schema as transient. Transient schemas do not have a Fail-safe period so they do not incur additional storage costs once they leave Time Travel; however, this means they are also not protected by Fail-safe in the event of a data loss. | true     | false     | false    | false   |
| name                | string | Specifies the identifier for the schema; must be unique for the database in which the schema is created.                                                                                                                                               | false    | true      | false    | <nil>   |

### snowflake_schema_grant

**Note**: The snowflake_schema_grant resource creates exclusive attachments of grants.
Across the entire Snowflake account, all of the schemas to which a single grant is attached must be declared
by a single snowflake_schema_grant resource. This means that even any snowflake_schema that have the attached
grant via any other mechanism (including other Terraform resources) will have that attached grant revoked by this resource.
These resources do not enforce exclusive attachment of a grant, it is the user's responsibility to enforce this.

#### properties

|     NAME      |  TYPE  |                                                                  DESCRIPTION                                                                  | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|---------------|--------|-----------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| database_name | string | The name of the database containing the schema on which to grant privileges.                                                                  | false    | true      | false    | <nil>   |
| privilege     | string | The privilege to grant on the schema.  Note that if "OWNERSHIP" is specified, ensure that the role that terraform is using is granted access. | true     | false     | false    | "USAGE" |
| roles         | set    | Grants privilege to these roles.                                                                                                              | true     | false     | false    | <nil>   |
| schema_name   | string | The name of the schema on which to grant privileges.                                                                                          | false    | true      | false    | <nil>   |
| shares        | set    | Grants privilege to these shares.                                                                                                             | true     | false     | false    | <nil>   |

### snowflake_share

#### properties

|   NAME   |  TYPE  |                                              DESCRIPTION                                              | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|----------|--------|-------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| accounts | list   | A list of accounts to be added to the share.                                                          | true     | false     | false    | <nil>   |
| comment  | string | Specifies a comment for the managed account.                                                          | true     | false     | false    | <nil>   |
| name     | string | Specifies the identifier for the share; must be unique for the account in which the share is created. | false    | true      | false    | <nil>   |

### snowflake_stage

#### properties

|        NAME         |  TYPE  |                                                                                     DESCRIPTION                                                                                     | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|---------------------|--------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| aws_external_id     | string |                                                                                                                                                                                     | true     | false     | true     | <nil>   |
| comment             | string | Specifies a comment for the stage.                                                                                                                                                  | true     | false     | false    | <nil>   |
| copy_options        | string | Specifies the copy options for the stage.                                                                                                                                           | true     | false     | false    | <nil>   |
| credentials         | string | Specifies the credentials for the stage.                                                                                                                                            | true     | false     | false    | <nil>   |
| database            | string | The database in which to create the stage.                                                                                                                                          | false    | true      | false    | <nil>   |
| encryption          | string | Specifies the encryption settings for the stage.                                                                                                                                    | true     | false     | false    | <nil>   |
| file_format         | string | Specifies the file format for the stage.                                                                                                                                            | true     | false     | false    | <nil>   |
| name                | string | Specifies the identifier for the stage; must be unique for the database and schema in which the stage is created.                                                                   | false    | true      | false    | <nil>   |
| schema              | string | The schema in which to create the stage.                                                                                                                                            | false    | true      | false    | <nil>   |
| snowflake_iam_user  | string |                                                                                                                                                                                     | true     | false     | true     | <nil>   |
| storage_integration | string | Specifies the name of the storage integration used to delegate authentication responsibility for external cloud storage to a Snowflake identity and access management (IAM) entity. | true     | false     | false    | <nil>   |
| url                 | string | Specifies the URL for the stage.                                                                                                                                                    | true     | false     | false    | <nil>   |

### snowflake_stage_grant

**Note**: The snowflake_stage_grant resource creates exclusive attachments of grants.
Across the entire Snowflake account, all of the stages to which a single grant is attached must be declared
by a single snowflake_stage_grant resource. This means that even any snowflake_stage that have the attached
grant via any other mechanism (including other Terraform resources) will have that attached grant revoked by this resource.
These resources do not enforce exclusive attachment of a grant, it is the user's responsibility to enforce this.

#### properties

|     NAME      |  TYPE  |                                     DESCRIPTION                                     | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|---------------|--------|-------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| database_name | string | The name of the database containing the current stage on which to grant privileges. | false    | true      | false    | <nil>   |
| privilege     | string | The privilege to grant on the stage.                                                | true     | false     | false    | "USAGE" |
| roles         | set    | Grants privilege to these roles.                                                    | true     | false     | false    | <nil>   |
| schema_name   | string | The name of the schema containing the current stage on which to grant privileges.   | false    | true      | false    | <nil>   |
| shares        | set    | Grants privilege to these shares.                                                   | true     | false     | false    | <nil>   |
| stage_name    | string | The name of the stage on which to grant privileges.                                 | false    | true      | false    | <nil>   |

### snowflake_storage_integration

#### properties

|           NAME            |  TYPE  |                                                  DESCRIPTION                                                  | OPTIONAL | REQUIRED  | COMPUTED |     DEFAULT      |
|---------------------------|--------|---------------------------------------------------------------------------------------------------------------|----------|-----------|----------|------------------|
| azure_tenant_id           | string |                                                                                                               | true     | false     | false    | ""               |
| comment                   | string |                                                                                                               | true     | false     | false    | ""               |
| created_on                | string | Date and time when the storage integration was created.                                                       | false    | false     | true     | <nil>            |
| enabled                   | bool   |                                                                                                               | true     | false     | false    | true             |
| name                      | string |                                                                                                               | false    | true      | false    | <nil>            |
| storage_allowed_locations | list   | Explicitly limits external stages that use the integration to reference one or more storage locations.        | false    | true      | false    | <nil>            |
| storage_aws_external_id   | string | The external ID that Snowflake will use when assuming the AWS role.                                           | false    | false     | true     | <nil>            |
| storage_aws_iam_user_arn  | string | The Snowflake user that will attempt to assume the AWS role.                                                  | false    | false     | true     | <nil>            |
| storage_aws_role_arn      | string |                                                                                                               | true     | false     | false    | ""               |
| storage_blocked_locations | list   | Explicitly prohibits external stages that use the integration from referencing one or more storage locations. | true     | false     | false    | <nil>            |
| storage_provider          | string |                                                                                                               | false    | true      | false    | <nil>            |
| type                      | string |                                                                                                               | true     | false     | false    | "EXTERNAL_STAGE" |

### snowflake_table_grant

**Note**: The snowflake_table_grant resource creates exclusive attachments of grants.
Across the entire Snowflake account, all of the tables to which a single grant is attached must be declared
by a single snowflake_table_grant resource. This means that even any snowflake_table that have the attached
grant via any other mechanism (including other Terraform resources) will have that attached grant revoked by this resource.
These resources do not enforce exclusive attachment of a grant, it is the user's responsibility to enforce this.

#### properties

|     NAME      |  TYPE  |                                                                           DESCRIPTION                                                                           | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT  |
|---------------|--------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|----------|
| database_name | string | The name of the database containing the current or future tables on which to grant privileges.                                                                  | false    | true      | false    | <nil>    |
| on_future     | bool   | When this is set to true, apply this grant on all future tables in the given schema.  The table_name and shares fields must be unset in order to use on_future. | true     | false     | false    | false    |
| privilege     | string | The privilege to grant on the current or future table.                                                                                                          | true     | false     | false    | "SELECT" |
| roles         | set    | Grants privilege to these roles.                                                                                                                                | true     | false     | false    | <nil>    |
| schema_name   | string | The name of the schema containing the current or future tables on which to grant privileges.                                                                    | true     | false     | false    | "PUBLIC" |
| shares        | set    | Grants privilege to these shares (only valid if on_future is unset).                                                                                            | true     | false     | false    | <nil>    |
| table_name    | string | The name of the table on which to grant privileges immediately (only valid if on_future is unset).                                                              | true     | false     | false    | <nil>    |

### snowflake_user

#### properties

|         NAME         |  TYPE  |                                                                                                        DESCRIPTION                                                                                                         | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|----------------------|--------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| comment              | string |                                                                                                                                                                                                                            | true     | false     | false    | <nil>   |
| default_namespace    | string | Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.                                                                                                | true     | false     | false    | <nil>   |
| default_role         | string | Specifies the role that is active by default for the user’s session upon login.                                                                                                                                            | true     | false     | true     | <nil>   |
| default_warehouse    | string | Specifies the virtual warehouse that is active by default for the user’s session upon login.                                                                                                                               | true     | false     | false    | <nil>   |
| disabled             | bool   |                                                                                                                                                                                                                            | true     | false     | true     | <nil>   |
| has_rsa_public_key   | bool   | Will be true if user as an RSA key set.                                                                                                                                                                                    | false    | false     | true     | <nil>   |
| login_name           | string | The name users use to log in. If not supplied, snowflake will use name instead.                                                                                                                                            | true     | false     | true     | <nil>   |
| must_change_password | bool   | Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.                                                                                         | true     | false     | false    | <nil>   |
| name                 | string | Name of the user. Note that if you do not supply login_name this will be used as login_name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)                              | false    | true      | false    | <nil>   |
| password             | string | **WARNING:** this will put the password in the terraform state file. Use carefully.                                                                                                                                        | true     | false     | false    | <nil>   |
| rsa_public_key       | string | Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.                                                                                                       | true     | false     | false    | <nil>   |
| rsa_public_key_2     | string | Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer. | true     | false     | false    | <nil>   |

### snowflake_view

#### properties

|   NAME    |  TYPE  |                                                          DESCRIPTION                                                          | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT  |
|-----------|--------|-------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|----------|
| comment   | string | Specifies a comment for the view.                                                                                             | true     | false     | false    | <nil>    |
| database  | string | The database in which to create the view. Don't use the | character.                                                          | false    | true      | false    | <nil>    |
| is_secure | bool   | Specifies that the view is secure.                                                                                            | true     | false     | false    | false    |
| name      | string | Specifies the identifier for the view; must be unique for the schema in which the view is created. Don't use the | character. | false    | true      | false    | <nil>    |
| schema    | string | The schema in which to create the view. Don't use the | character.                                                            | true     | false     | false    | "PUBLIC" |
| statement | string | Specifies the query used to create the view.                                                                                  | false    | true      | false    | <nil>    |

### snowflake_view_grant

**Note**: The snowflake_view_grant resource creates exclusive attachments of grants.
Across the entire Snowflake account, all of the views to which a single grant is attached must be declared
by a single snowflake_view_grant resource. This means that even any snowflake_view that have the attached
grant via any other mechanism (including other Terraform resources) will have that attached grant revoked by this resource.
These resources do not enforce exclusive attachment of a grant, it is the user's responsibility to enforce this.

#### properties

|     NAME      |  TYPE  |                                                                          DESCRIPTION                                                                          | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT  |
|---------------|--------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|----------|
| database_name | string | The name of the database containing the current or future views on which to grant privileges.                                                                 | false    | true      | false    | <nil>    |
| on_future     | bool   | When this is set to true, apply this grant on all future views in the given schema.  The view_name and shares fields must be unset in order to use on_future. | true     | false     | false    | false    |
| privilege     | string | The privilege to grant on the current or future view.                                                                                                         | true     | false     | false    | "SELECT" |
| roles         | set    | Grants privilege to these roles.                                                                                                                              | true     | false     | false    | <nil>    |
| schema_name   | string | The name of the schema containing the current or future views on which to grant privileges.                                                                   | true     | false     | false    | "PUBLIC" |
| shares        | set    | Grants privilege to these shares (only valid if on_future is unset).                                                                                          | true     | false     | false    | <nil>    |
| view_name     | string | The name of the view on which to grant privileges immediately (only valid if on_future is unset).                                                             | true     | false     | false    | <nil>    |

### snowflake_warehouse

#### properties

|             NAME             |  TYPE  |                                                               DESCRIPTION                                                                | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|------------------------------|--------|------------------------------------------------------------------------------------------------------------------------------------------|----------|-----------|----------|---------|
| auto_resume                  | bool   | Specifies whether to automatically resume a warehouse when a SQL statement (e.g. query) is submitted to it.                              | true     | false     | true     | <nil>   |
| auto_suspend                 | int    | Specifies the number of seconds of inactivity after which a warehouse is automatically suspended.                                        | true     | false     | true     | <nil>   |
| comment                      | string |                                                                                                                                          | true     | false     | false    | ""      |
| initially_suspended          | bool   | Specifies whether the warehouse is created initially in the ‘Suspended’ state.                                                           | true     | false     | false    | <nil>   |
| max_cluster_count            | int    | Specifies the maximum number of server clusters for the warehouse.                                                                       | true     | false     | true     | <nil>   |
| min_cluster_count            | int    | Specifies the minimum number of server clusters for the warehouse (only applies to multi-cluster warehouses).                            | true     | false     | true     | <nil>   |
| name                         | string |                                                                                                                                          | false    | true      | false    | <nil>   |
| resource_monitor             | string | Specifies the name of a resource monitor that is explicitly assigned to the warehouse.                                                   | true     | false     | true     | <nil>   |
| scaling_policy               | string | Specifies the policy for automatically starting and shutting down clusters in a multi-cluster warehouse running in Auto-scale mode.      | true     | false     | true     | <nil>   |
| statement_timeout_in_seconds | int    | Specifies the time, in seconds, after which a running SQL statement (query, DDL, DML, etc.) is canceled by the system                    | true     | false     | false    |       0 |
| wait_for_provisioning        | bool   | Specifies whether the warehouse, after being resized, waits for all the servers to provision before executing any queued or new queries. | true     | false     | false    | <nil>   |
| warehouse_size               | string |                                                                                                                                          | true     | false     | true     | <nil>   |

### snowflake_warehouse_grant

**Note**: The snowflake_warehouse_grant resource creates exclusive attachments of grants.
Across the entire Snowflake account, all of the warehouses to which a single grant is attached must be declared
by a single snowflake_warehouse_grant resource. This means that even any snowflake_warehouse that have the attached
grant via any other mechanism (including other Terraform resources) will have that attached grant revoked by this resource.
These resources do not enforce exclusive attachment of a grant, it is the user's responsibility to enforce this.

#### properties

|      NAME      |  TYPE  |                       DESCRIPTION                       | OPTIONAL | REQUIRED  | COMPUTED | DEFAULT |
|----------------|--------|---------------------------------------------------------|----------|-----------|----------|---------|
| privilege      | string | The privilege to grant on the warehouse.                | true     | false     | false    | "USAGE" |
| roles          | set    | Grants privilege to these roles.                        | true     | false     | false    | <nil>   |
| warehouse_name | string | The name of the warehouse on which to grant privileges. | false    | true      | false    | <nil>   |
<!-- END -->

## Development

To do development you need Go installed, this repo cloned and that's about it.
It has not been tested on Windows, so if you find problems let us know.

<<<<<<< HEAD
If you want to build and test the provider localling there is a make target
`make install-tf` that will build the provider binary and install it in a
location that terraform can find.
=======
If you want to build and test the provider locally there is a make target `make install-tf` that will build the provider binary and install it in a location that terraform can find.
>>>>>>> 4346dee4430764fa7c608a761b5ffd3c53650631

## Testing

For the Terraform resources, there are 3 levels of testing - internal, unit and
acceptance tests.

The 'internal' tests are run in the
`github.com/viostream/terraform-provider-snowflake/pkg/resources` package so
that they can test functions that are not exported. These tests are intended to
be limited to unit tests for simple functions.

The 'unit' tests are run in
`github.com/viostream/terraform-provider-snowflake/pkg/resources_test`, so
they only have access to the exported methods of `resources`. These tests
exercise the CRUD methods that on the terraform resources. Note that all tests
here make use of database mocking and are run locally. This means the tests are
fast, but are liable to be wrong in suble ways (since the mocks are unlikely to
be perfect).

You can run these first two sets of tests with `make test`.

The 'acceptance' tests run the full stack, creating, modifying and destroying
resources in a live snowflake account. To run them you need a snowflake account
and the proper environment variables set: SNOWFLAKE_ACCOUNT, SNOWFLAKE_USER,
SNOWFLAKE_PASSWORD, SNOWFLAKE_ROLE. These tests are slower but have higher
fidelity.

To run all tests, including the acceptance tests, run `make test-acceptance`.

<<<<<<< HEAD
Note that we also run all tests in our [Travis-CI account](
https://travis-ci.com/viostream/terraform-provider-snowflake).
=======
We also run all tests in our [Travis-CI account](https://travis-ci.com/chanzuckerberg/terraform-provider-snowflake).

### Pull Request CI

Our CI jobs run the full acceptence test suite, which involves creating and destroying resources in a live snowflake account. Travis-CI is configured with environment variables to authenticate to our test snowflake account. For security reasons, those variables are not available to forks of this repo.

If you are making a PR from a forked repo, you can create a new Snowflake trial account and set up Travis to build it by setting these environement variables:

* `SNOWFLAKE_ACCOUNT` - The account name
* `SNOWFLAKE_USER` - A snowflake user for running tests.
* `SNOWFLAKE_PASSWORD` - Password for that user.
* `SNOWFLAKE_ROLE` - Needs to be ACCOUNTADMIN or similar.
* `SNOWFLAKE_REGION` - Default is us-west-2, set this if your snowflake account is in a different region.

If you are using the Standard Snowflake plan, it's recommended you also set up the following environment variables to skip tests for features not enabled for it:

* `SKIP_DATABASE_TESTS` - to skip tests with retention time larger than 1 day
* `SKIP_WAREHOUSE_TESTS` - to skip tests with multi warehouses

>>>>>>> 4346dee4430764fa7c608a761b5ffd3c53650631

# Connect-DB-Demo
A demo dispaly how to connect to ibm cloud database instance via golang.

I prefer to use IBM [CLI](https://github.com/IBM-Cloud/ibm-cloud-cli-release/releases/) to demo

### Provision one instance 
> skip this step if you already have one

1. To create a database, you’ll need to log in first. If you’re using a federated identity, you’ll use:
```
ibmcloud login --sso
```
2. Create your databases
```
ibmcloud resource service-instance-create <instance_name> databases-for-postgresql standard us-south
```
for example:
```
ibmcloud resource service-instance-create 'Databases for PostgreSQL dev-yp-02-zw' databases-for-postgresql standard us-south
```
1. Once the database is done provisioning, you can get the credentials using the cloud-databases plugin. Install that using the following command
```
ibmcloud plugin install cloud-databases
```
4. Get ```hostname```, ```port```, ```sslmode```, ```database``` via cli and set them as environment variables,
```
ibmcloud cdb cxn 'Databases for PostgreSQL dev-yp-02-zw'

Retrieving public connection strings for Databases for PostgreSQL dev-yp-02-zw...
OK

Type         Connection String
PostgreSQL   postgres://admin:$PASSWORD@3413a99c-c92a-46aa-892a-b53cc0b86c91.f4fe0c06b8ef4fffb4ab4460b0f2ad58.databases.appdomain.cloud:30816/ibmclouddb?sslmode=verify-full
CLI          PGPASSWORD=$PASSWORD PGSSLROOTCERT=03b6f19d-a31a-11e9-865e-dee1fee4feba psql 'host=3413a99c-c92a-46aa-892a-b53cc0b86c91.f4fe0c06b8ef4fffb4ab4460b0f2ad58.databases.appdomain.cloud port=30816 dbname=ibmclouddb user=admin sslmode=verify-full'
```
```
export HOSTNAME='3413a99c-c92a-46aa-892a-b53cc0b86c91.f4fe0c06b8ef4fffb4ab4460b0f2ad58.databases.appdomain.cloud'
export PORT=30816
export DATABASE=ibmclouddb
export SSLMODE='verify-full'
```
5. Set password for databases
```
ibmcloud cdb deployment-user-password 'Databases for PostgreSQL dev-yp-02-zw' admin your_password
The user's password is being changed with this task:

Key                   Value
ID                    crn:v1:bluemix:public:databases-for-postgresql-dev-yp-02:us-south:a/b2742234a9bf412a8183644a5a92cd95:3413a99c-c92a-46aa-892a-b53cc0b86c91:task:7ec562b5-da52-4b05-bc23-370ba2dfe085
Deployment ID         crn:v1:bluemix:public:databases-for-postgresql-dev-yp-02:us-south:a/b2742234a9bf412a8183644a5a92cd95:3413a99c-c92a-46aa-892a-b53cc0b86c91::
Description           Updating user.
Created At            2019-09-05T06:52:04Z
Status                running
Progress Percentage   0

Status                completed
Progress Percentage   100
Location              https://api.dev-yp-02.us-south.databases.cloud.ibm.com/v4/ibm/deployments/crn:v1:bluemix:public:databases-for-postgresql-dev-yp-02:us-south:a%2Fb2742234a9bf412a8183644a5a92cd95:3413a99c-c92a-46aa-892a-b53cc0b86c91::
OK
```
```
export USERNAME=admin
export PASSWORD=<your_password>
```
6. You’ll also need to decode the CA certificate that your databases need for authentication. To decode it, run the following command then make sure to copy the decoded certificate and save it to a file on your system
```
ibmcloud cdb cacert 'Databases for PostgreSQL dev-yp-02-zw'
```
```
export CAFILE=your_cafile_path
```
7. Run it, and it will show all tables in this db
```
make run
```
```
demo
pg_statistic
pg_type
pg_policy
pg_authid
pg_user_mapping
pg_subscription
pg_attribute
pg_proc
pg_class
pg_attrdef
pg_constraint
pg_inherits
pg_index
pg_operator
pg_opfamily
pg_opclass
pg_am
pg_amop
pg_amproc
pg_language
pg_largeobject_metadata
pg_aggregate
pg_statistic_ext
pg_rewrite
pg_trigger
pg_event_trigger
pg_description
pg_cast
pg_enum
pg_namespace
pg_conversion
pg_depend
pg_database
pg_db_role_setting
pg_tablespace
pg_pltemplate
pg_auth_members
pg_shdepend
pg_shdescription
pg_ts_config
pg_ts_config_map
pg_ts_dict
pg_ts_parser
pg_ts_template
pg_extension
pg_foreign_data_wrapper
pg_foreign_server
pg_foreign_table
pg_replication_origin
pg_default_acl
pg_init_privs
pg_seclabel
pg_shseclabel
pg_collation
pg_partitioned_table
pg_range
pg_transform
pg_sequence
pg_publication
pg_publication_rel
pg_subscription_rel
pg_largeobject
sql_parts
sql_languages
sql_features
sql_implementation_info
sql_packages
sql_sizing
sql_sizing_profiles
```
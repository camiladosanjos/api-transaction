create table "Accounts" (
	"Account_ID" serial primary key,
	"Document_Number" varchar not null
);

create table "OperationsTypes"(
	"OperationType_ID" integer primary key,
	"Description" varchar not null,
	"Sign" smallint not null
);

insert into "OperationsTypes" values (1, 'COMPRA A VISTA', -1),
(2, 'COMPRA PARCELADA', -1), (3, 'SAQUE', -1), (4, 'PAGAMENTO', 1);

create table "Transactions" (
	"Transaction_ID" serial primary key,
	"Account_ID" integer not null,
	"OperationType_ID" integer not null,
	"Amount" decimal not null,
	"EventDate" timestamp not null,
	foreign key ("Account_ID") references "Accounts" ("Account_ID"),
	foreign key ("OperationType_ID") references "OperationsTypes" ("OperationType_ID")
);

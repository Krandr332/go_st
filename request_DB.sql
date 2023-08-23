CREATE TABLE "public.user" (
	"user_id" serial NOT NULL,
	"username" varchar(255) NOT NULL,
	"email" varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	CONSTRAINT "user_pk" PRIMARY KEY ("user_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.friends" (
	"friendship_id" serial NOT NULL,
	"user_id1" int NOT NULL,
	"user_id2" int NOT NULL,
	"status" varchar(50) NOT NULL,
	CONSTRAINT "friends_pk" PRIMARY KEY ("friendship_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.post" (
	"post_id" serial NOT NULL,
	"user_id" int NOT NULL,
	"photo_url" varchar(255) NOT NULL,
	"text" TEXT NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT 'now',
	CONSTRAINT "post_pk" PRIMARY KEY ("post_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.message" (
	"message_id" serial NOT NULL,
	"sender_id" int NOT NULL,
	"recipient_id" int NOT NULL,
	"content" TEXT NOT NULL,
	"sent_at" TIMESTAMP NOT NULL DEFAULT 'now',
	CONSTRAINT "message_pk" PRIMARY KEY ("message_id")
) WITH (
  OIDS=FALSE
);




ALTER TABLE "public.friends" ADD CONSTRAINT "friends_fk0" FOREIGN KEY ("user_id1") REFERENCES "public.user"("user_id");

ALTER TABLE "public.post" ADD CONSTRAINT "post_fk0" FOREIGN KEY ("user_id") REFERENCES "public.user"("user_id");

ALTER TABLE "public.message" ADD CONSTRAINT "message_fk0" FOREIGN KEY ("sender_id") REFERENCES "public.user"("user_id");






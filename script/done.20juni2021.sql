
ALTER TABLE public.po ADD sales_id int8 NULL DEFAULT 0;

update  m_menus set link = 'purchase-order' where  "name" = 'po'

ALTER TABLE public.sales_order ADD is_cash bool null default false;

INSERT INTO public.customer
(id, "name", code, top, address1, address2, address3, address4, contact_person, phone_number, status, last_update_by, last_update)
VALUES(99999999, 'Customer Cash', '999999', 0, 'Cash', 'Cash', 'Cash', 'Cash', 'Cash', 'Cash', 1, 'system', CURRENT_DATE);



CREATE TABLE public.parameter (
	id bigserial NOT NULL,
	"name" varchar(255) NULL,
	value varchar(255) NULL,
	isviewable int2 NULL DEFAULT 1,
	PRIMARY KEY (id)
);

INSERT INTO public."parameter" ("name", value, isviewable) values ('tax', '10', 1);

alter table po add column is_tax bool default false;

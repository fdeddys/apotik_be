INSERT INTO public.m_menus
(name, description, link, parent_id, status, icon, "ordering")
VALUES( 'payment-supplier', 'Payment to Supplier', 'payment-supplier', 
(select id from m_menus mm where mm."name" ='purchasing')
, 1, '', '240');



CREATE TABLE payment_supplier (
	id bigserial NOT NULL,
	payment_no varchar(255) NULL,
	payment_date timestamptz NULL,
	supplier_id int8 NULL DEFAULT 0,
	payment_type_id int8 NULL DEFAULT 0,
	payment_reff varchar NULL DEFAULT 0,
	note text NULL,
	total  numeric NULL DEFAULT 0,
	status int2 NULL DEFAULT 1,
	last_update_by varchar(255) NULL,
	last_update timestamptz NULL,
	PRIMARY KEY (id)
);


CREATE TABLE payment_supplier_detail (
	id bigserial NOT NULL,
	payment_supplier_id int8 default 0,
	receiving_id int8 NULL DEFAULT 0,
	total numeric NULL DEFAULT 0,
	last_update_by varchar(255) NULL,
	last_update timestamptz NULL,
	PRIMARY KEY (id)
);


ALTER TABLE public.receive ADD is_paid boolean NULL DEFAULT false;

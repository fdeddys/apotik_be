alter table mutation rename column totat to total;

ALTER TABLE public.product ADD hpp numeric default 0;

alter table history_stock 
add column warehouse_id int8 NULL DEFAULT 0

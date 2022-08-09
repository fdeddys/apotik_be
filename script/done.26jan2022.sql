INSERT INTO public.m_menus
( "name", description, link, parent_id, status, icon, "ordering")
VALUES('report-payment-supplier', 'Report Payment Supplier', 'report-payment-supplier', 
(select id from m_menus m where m."name" ='report')
, 1, '', '640');


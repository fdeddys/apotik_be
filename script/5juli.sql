
INSERT INTO public.m_menus 
("name",description,link,parent_id,status,icon,"ordering") 
VALUES
('direct-sales-payment','Direct Sales Payment','direct-sales-payment',15,1,'','90');


INSERT INTO public.m_role_menu 
(role_id,menu_id,status,last_update_by,last_update) 
VALUES
(1,
(select id from m_menus where name = 'direct-sales-payment')
,1,'system',CURRENT_TIMESTAMP);

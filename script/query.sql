19 juli

INSERT INTO public.m_menus(
	 name, description, link, parent_id, status, icon, ordering)
	VALUES ( 
		'parameter', 'Parameter', 'apotik-param', 6, 1, '', 650);
	
INSERT into m_role_menu(
	 role_id, menu_id, status, last_update_by, last_update)
	 	
	select 1, id, 1, 'system', CURRENT_DATE from m_menus where name='parameter'

	
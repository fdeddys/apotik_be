package database

import "distribution-system-be/models/dto"

func ReportPaymentCashByDate(dateStart, dateEnd string) []dto.ReportPaymentCash {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var datas []dto.ReportPaymentCash

	db.Raw("select l.name payment_type_name, p.payment_no , to_char(p.payment_date, 'DD/Mon/YYYY') payment_date, "+
		" so.sales_order_no , to_char(so.order_date , 'DD/Mon/YYYY') order_date, p.total_order , p.total_payment , "+
		" to_char(p.last_update, 'DD/Mon/YYYY') last_update, p.last_update_by "+
		" from payment p "+
		" inner join payment_detail pd on pd.payment_id = p.id and p.is_cash = true and p.status = 20 "+
		" left join lookup l on l.id = pd.payment_type_id  "+
		" inner join payment_order po on po.payment_id  = p.id "+
		" inner join sales_order so on po.sales_order_id = so.id  "+
		" where payment_date between ?  and ? ", dateStart, dateEnd).Scan(&datas)

	return datas

}

func ReportSalesByDate(dateStart, dateEnd string) []dto.ReportSales {
	db := GetDbCon()
	db.Debug().LogMode(true)

	var datas []dto.ReportSales

	db.Raw("select "+
		" to_char(so.order_date , 'DD/Mon/YYYY') as order_date , so.sales_order_no , so.status , "+
		" p.plu, p.name as product_name , sod.qty_order , "+
		" l.name as uom , floor(sod.price) , sod.disc1  "+
		" from sales_order so  "+
		" inner join sales_order_detail sod on sod.sales_order_id = so.id "+
		" left join product p on sod.product_id = p.id "+
		" left join lookup l on l.id = sod.uom  "+
		" where so.order_date between ?  and ? ", dateStart, dateEnd).Scan(&datas)

	return datas

}

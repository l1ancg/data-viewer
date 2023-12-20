package utils

import (
	"strings"
	"testing"

	"github.com/xwb1989/sqlparser"
)

func Test_Parse(t *testing.T) {
	stmt, err := sqlparser.Parse("SELECT `tos`.`transport_number` AS `transport_number`, `toto`.`id` AS `id`, `tos`.`id` AS `id`, `tosd`.`id` AS `id`, `tosd`.`status` AS `status`, `tosdi`.`id` AS `id`, `tosdi`.`status` AS `status`, `tosdi`.`amount` AS `amount`, `tosdi`.`income_uid` AS `income_uid`, `tosdi`.`income_user_name` AS `income_user_name`, `tosdif`.`id` AS `id`, `tosdif`.`status` AS `status`, `omp`.`id` AS `id`, `omp`.`status` AS `status`, `toss`.`id` AS `id`, `tossp`.`id` AS `id`, `tossp`.`status` AS `status`, `tossp`.`spending_uid` AS `spending_uid`, `tossp`.`spending_user_name` AS `spending_user_name`, `tosa`.`id` AS `id`, `tosai`.`id` AS `id`, `tosai`.`status` AS `status`, `tosai`.`merge_pay_id` AS `merge_pay_id`, `omp2`.`status` AS `status`, `tosai`.`income_user_master_id` AS `income_user_master_id`, `tosai`.`income_user_name` AS `income_user_name`, `tosai`.`type` AS `type`, `tosai`.`amount` AS `amount`, `tos`.`invoice_enterprise_id` AS `invoice_enterprise_id`, `tos`.`proxy_invoice_enterprise_id` AS `proxy_invoice_enterprise_id` FROM `t_oas_settlement` AS `tos` LEFT JOIN `t_oas_settlement_driver` AS `tosd` ON `tosd`.`settlement_id` = `tos`.`id` LEFT JOIN `t_oas_settlement_driver_income` AS `tosdi` ON `tosdi`.`settlement_driver_id` = `tosd`.`id` LEFT JOIN `t_oas_settlement_driver_income_fuel` AS `tosdif` ON `tosdif`.`settlement_driver_id` = `tosd`.`id` LEFT JOIN `t_oas_merge_pay` AS `omp` ON `tosdi`.`merge_pay_id` = `omp`.`id` LEFT JOIN `t_oas_settlement_shipper` AS `toss` ON `toss`.`settlement_id` = `tosd`.`settlement_id` LEFT JOIN `t_oas_settlement_shipper_spending` AS `tossp` ON `tossp`.`settlement_shipper_id` = `toss`.`id` AND `tosdi`.`type` = `tossp`.`type` LEFT JOIN `t_oas_settlement_agent` AS `tosa` ON `tosa`.`settlement_id` = `tos`.`id` LEFT JOIN `t_oas_settlement_agent_income` AS `tosai` ON `tosai`.`settlement_agent_id` = `tosa`.`id` LEFT JOIN `t_oas_merge_pay` AS `omp2` ON `tosai`.`merge_pay_id` = `omp2`.`id` INNER JOIN `t_oas_transport_order` AS `toto` ON `toto`.`id` = `tos`.`transport_order_id` AND `toto`.`third_party_payment_platform` = 0 WHERE `tos`.`type` = ? AND `tos`.`status` = ? AND `tos`.`id` IN (SELECT `id` FROM `table1` WHERE `type` = ?)")
	if err != nil {
		t.Log("Parse error: ", err.Error())
		return
	}

	// Otherwise do something with stmt
	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		// 创建一个新的 AndExpr 来保存修改后的条件
		newAndExpr := &sqlparser.AndExpr{}
		// 遍历现有的条件
		switch node := stmt.Where.Expr.(type) {
		case *sqlparser.AndExpr:
			// 如果找到要删除的条件，则不添加到新的 AndExpr 中
			if !strings.Contains(sqlparser.String(node.Left), "tos.transport_number") {
				newAndExpr.Left = node.Left
			}
			if !strings.Contains(sqlparser.String(node.Right), "tos.transport_number") {
				newAndExpr.Right = node.Right
			}
		}
		t.Log(sqlparser.String(stmt))
	}
}

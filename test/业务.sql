恢复数据：

UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1166, 1167, 1168) and created=0 and status = 1 and event_type= 2;
UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1166, 1167, 1168) and created=0;
UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1169, 1170, 1171) and created=0;


UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1172,1173,1174,1199) and created=0 and status = 1 and event_type = 2;
UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1175,1176,1177,1200) and created=0 and status = 1 and event_type = 2;
UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1178,1179,1180) and created=0 and status = 1 and event_type = 2;
UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1181,1182,1183) and created=0 and status = 1 and event_type = 2;
UPDATE `task_log` SET `created` = '1687742068' WHERE `task_id` in (1172,1173,1174,1199,1175,1176,1177,1200,1178,1179,1180,1181,1182,1183) and created=0 and status = 1 and event_type = 2;








CREATE TABLE `data_statistics` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '记录id',
  `task_id` varchar(200)  NOT NULL DEFAULT '0' COMMENT '任务id',
  `push_success_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'push的成功数',
  `push_topic_success_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'hokuto_msg_topic的push的成功数',
  `push_report_click_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'push的点击数',
  `push_report_reach_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'push的到达数',
  `push_order_loan_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'push的用户的下单数',
  `status` tinyint(3) unsigned NOT NULL DEFAULT 1 COMMENT '数据状态: 2:无效,1:有效',
  `created` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '任务创建时间',
  `updated` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '任务修改时间',
  PRIMARY KEY (`id`),
  KEY `IDX_created` (`created`),
  KEY `IDX_task_id` (`task_id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='数据统计表';




我的验证逻辑

1.1 luban.task_log记录表
 a)、获取被过滤的用户的uid
 b)、获取成功提额的用户uid
 1.2 crm_feature.rc_outdata_repayment_result校验数据
 a)、校验被过滤的用户的uid
 b)、校验被成功提额的用户的uid的数据



smsContent = strings.Replace(smsContent, "{{amount}}", fmt.Sprint(Amount), -1)
smsContent = strings.Replace(smsContent, "{{new_amount}}", fmt.Sprint(result["new_amount"]), -1)
pushContent = strings.Replace(pushContent, "{{amount}}", fmt.Sprint(Amount), -1)
pushContent = strings.Replace(pushContent, "{{new_amount}}", fmt.Sprint(result["new_amount"]), -1)

2023年06月25日 星期日
select count(id) as num ,event_type,task_id from task_log where task_id in (1178,1179,1180,1181,1182,1183) and status = 1  and created >= 1687648480 group by  event_type,  task_id 
select count(id) as num ,event_type,task_id from task_log where task_id in (1172,1173,1174,1175,1176,1177,1199,1200) and status = 1  and created >= 1687648480 group by  event_type,  task_id 
select count(id) as num ,event_type,task_id from task_log where task_id in (1166,1185,1167,1172,1178,1173,1199,1179) and status = 1  and created >= 1687648480 group by  event_type,  task_id 


(11890019,
11911001,
11235585,
11911707,
11904781,
11912083,
11913391,
11897445,
11905002,
11910503,
11897618,
11911304,
11912905,
11909400,
11911505,
9889486,
11315932)


SELECT t1.id,
         t1.uid,
         t1.ctime,
         t1.mtime ,
         substr(t1.ctime,
         1,
         10) AS c_date ,
         substr(t1.mtime,
         1,
         10) AS m_date ,
         JSON_EXTRACT(t1.data,
         "$.exter_first_trans_to_apply_days") AS exter_first_trans_to_apply_days ,
         JSON_EXTRACT(t1.data,
         "$.dz_level_short_v2") AS dz_level_short_v2
FROM crm_feature.rc_outdata_repayment_result AS t1
WHERE t1.uid IN(392777, 557693, 453657, 478583, 735299, 568351, 498486, 767183, 664355, 542291, 523762, 448753, 650480, 709578, 675660, 844780, 752175, 724967, 869497, 1136777, 630060) 



SELECT DISTINCT a.uid,
         c.product_id AS pid,
         d.total_amount AS amount
FROM 
    (SELECT uid,
         MAX(real_repayment_time) AS rtime
    FROM user_loan_plan
    WHERE real_repayment_time
        BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 4 DAY))
            AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 3 DAY))and `status`=2
    GROUP BY  uid) a
LEFT JOIN user_loan_plan b
    ON a.uid=b.uid
        AND b.ctime>a.rtime
LEFT JOIN user_product c
    ON a.uid=c.uid
LEFT JOIN product_account d
    ON a.uid=d.uid
WHERE c.is_pass='YES'
        AND c.product_id IN(108,119,116,120,121)
        AND b.id is NULL
        AND d.account_type=2
        AND a.uid % 100 >= 50
GROUP BY  a.uid

func (obj *uangmeRepo) DealData(ctx context.Context, taskId string, startTime int64, endTime int64) (bool, error) {
	num, err_task_log := obj.lubanDb.Master().Table("task_log").Where("created BETWEEN ? AND ?", startTime, endTime).In("task_id", strArr).FindAndCount(ctx, dataTaskLog1)
	win, err = obj.lubanDb.Master().Table("task_log").Select("batch_id").Where("created BETWEEN ? AND ?", startTime, endTime).In("task_id", strArr).Get(ctx, dataTaskLog2)
	err = obj.lubanDb.Master().Table("task_log").Select("batch_id").Where("created BETWEEN ? AND ?", startTime, endTime).In("task_id", strArr).GroupBy("batch_id").Find(ctx, data2)
	num2, errs1 := obj.cmpushInDb.Master().Table("hokuto_msg_topic").Where("state = ? AND success =  ?", 1, 1).In("push_id", strArr2).FindAndCount(ctx, hokutoMsgTopicData2)
	num3, errs2 := obj.cmpushInDb.Master().Table("hokuto_msg_report").Where("status >= ", 0).In("push_id", strArr2).FindAndCount(ctx, hokutoMsgTopicData2)
	num4, errs3 := obj.cmpushInDb.Master().Table("hokuto_msg_report").Where("status >= ", 1).In("push_id", strArr2).FindAndCount(ctx, hokutoMsgTopicData2)
	affect1, err1 := obj.db.Master().Table("data_statistics").Where("created = ? AND task_id = ?", dataStatics.Created, dataStatics.TaskId).Update(ctx, dataStaticsUpdate)
}










for i, c := range ups {
		switch c.ProductId {
		case 119:
			if i > 0 {
				remove := new(model.UserProduct)
				remove = ups[0]
				newUpss := make([]*model.UserProduct, 0)
				newUpss = append(ups[:0], ups[(0+1):]...)
				newUpss[0] = ups[i]
				newUpss = append(newUpss, remove)
				ups = newUpss
			}
		default:
			break
		}
	}

  {"sms_channel_type":"0","sms_content":"yugjgku"}





  SELECT DISTINCT a.uid,
         c.product_id AS pid,
         d.total_amount AS amount FROM
    (SELECT uid,
         MAX(real_repayment_time) AS rtime
    FROM user_loan_plan
    WHERE real_repayment_time
        BETWEEN UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 4 DAY))
            AND UNIX_TIMESTAMP(DATE_SUB(CURRENT_DATE,INTERVAL 3 DAY))and `status`=2
    GROUP BY  uid) a
LEFT JOIN user_loan_plan b
    ON a.uid=b.uid
        AND b.ctime>a.rtime
LEFT JOIN user_product c
    ON a.uid=c.uid
LEFT JOIN product_account d
    ON a.uid=d.uid
WHERE c.is_pass='YES'
        AND c.product_id IN(108,119,116,120,121)
        AND b.id is NULL
        AND d.account_type=2
        AND a.uid % 100 >= 50
GROUP BY  a.uid




template:
    config: '{
  "uangme": {
    "filters": [
      {
        "label": "黑名单过滤器",
        "name": "blacklist_filter",
        "params": {
          "filter": {
            "label": "过滤",
            "field_type": "input",
            "hidden": "hidden",
            "type": "int",
            "value": "1"
          }
        }
      },
      {
        "label": "用户短账龄等级过滤器",
        "name": "user_dz_level_short_v2_filter",
        "params": {
          "filter": {
            "label": "过滤",
            "field_type": "input",
            "type": "int",
            "value": "1"
          }
        }
      },
       {
            "label": "用户风控评分过滤器",
            "name": "user_risk_score_filter",
            "params": {
                "score": {
                    "label": "最低分值",
                    "field_type": "input",
                    "type": "int",
                    "value": "620"
                }
            }
        },
        {
            "label": "用户账龄过滤器",
            "name": "user_account_age_filter",
            "params": {
                "condition": {
                    "label": "条件",
                    "field_type": "input",
                    "type": "",
                    "value": ">="
                },
                "day": {
                    "label": "账龄(天)",
                    "field_type": "input",
                    "type": "int",
                    "value": "120"
                }
            }
        },
      {
            "label": "谷歌账号过滤器",
            "name": "google_account_filter",
            "params": {
                "uids": {
                    "label": "uids(逗号分隔)",
                    "field_type": "input",
                    "type": "",
                    "value": "6894787"
                }
            }
      },
      {
            "label": "重复提额限制过滤器",
            "name": "raise_amount_days_filter",
            "params": {
                "method": {
                    "label": "操作类型(逗号分隔)",
                    "field_type": "input",
                    "type": "",
                    "value": "recall"
                },
                "op": {
                    "label": "提降",
                    "field_type": "input",
                    "type": "",
                    "value": "add"
                },
                "day": {
                    "label": "天数",
                    "field_type": "input",
                    "type": "int",
                    "value": "18"
                }
            }
       },
       {
            "label": "逾期中过滤器",
            "name": "is_overdue_user_filter",
            "params": {
                "filter": {
                    "label": "过滤",
                    "field_type": "input",
                    "hidden": "hidden",
                    "type": "int",
                    "value": "1"
                }
            }
        },
      {
        "label": "批次过滤器",
        "name": "task_batch_filter",
        "params": {
          "batch_ids": {
            "label": "批次id(英文,逗号分隔)",
            "field_type": "input",
            "type": "",
            "value": ""
          }
        }
      },
      {
        "label": "风控分过滤器",
        "name": "rc_score_filter",
        "params": {
          "min_score": {
            "label": "最低分",
            "field_type": "input",
            "type": "int",
            "value": ""
          },
          "max_score": {
            "label": "最高分",
            "field_type": "input",
            "type": "int",
            "value": ""
          }
        }
      },
      {
        "label": "用户在贷过滤器",
        "name": "user_in_loan_filter",
        "params": {
          "filter": {
            "label": "过滤",
            "field_type": "input",
            "hidden": "hidden",
            "type": "int",
            "value": "1"
          }
        }
      }
    ],
    "actions": [
        {
            "label": "策略调额营销执行器",
            "name": "strategy_account_amount_action",
            "params": {
                "scene": {
                    "label": "场景",
                    "field_type": "select",
                    "options": [
                        {
                            "label": "贷前未借款",
                            "value": "noloan"
                        },
                        {
                            "label": "贷中",
                            "value": "inloan"
                        },
                        {
                            "label": "贷中v2",
                            "value": "inloanv2"
                        }
                    ],
                    "type": "",
                    "value": "noloan"
                },
                "amount": {
                    "label": "用户当前额度",
                    "field_type": "tag_input",
                    "type": "",
                    "value": ""
                },
                 "level": {
                    "label": "用户当前等级",
                    "field_type": "tag_input",
                    "type": "",
                    "value": "0"
                },
                "pid": {
                    "label": "授信通过的产品",
                    "field_type": "tag_input",
                    "type": "",
                    "value": ""
                },
                "is_change_user_account_amount": {
                    "label": "是否开启提额",
                    "field_type": "tag_input",
                    "type": "",
                    "value": ""
                },
                "sms_channel_type": {
                    "label": "短信通道",
                    "field_type": "select",
                    "options": [
                        {"label": "默认","value":"0"},
                        {"label": "召回营销","value":"6"},
                        {"label": "VIP通道","value":"9"}
                    ],
                    "type": "",
                    "value": "6"
                },
                "sms_content": {
                    "label": "短信内容",
                    "field_type": "tag_input",
                    "type": "",
                    "value": ""
                },
                "push_title": {
                    "label": "push标题",
                    "field_type": "tag_input",
                    "type": "",
                    "value": ""
                },
                "push_content": {
                    "label": "push内容",
                    "field_type": "tag_input",
                    "type": "",
                    "value": ""
                },
                "click_action": {
                    "label": "push跳转路由",
                    "field_type": "input",
                    "type": "",
                    "value": "uangme://api.uangme.com/main"
                }
            }
        },
      {
        "label": "发送短信",
        "name": "sendmsg_action",
        "params": {
          "sms_channel_type": {
            "label": "短信通道",
            "field_type": "select",
            "options": [
                {"label": "召回营销","value":"6"},
                {"label": "默认","value":"0"},
                {"label": "VIP通道","value":"9"}
            ],
            "type": "",
            "value": "6"
          },
          "sms_content": {
            "label": "短信内容",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          }
        }
      },
      {
        "label": "发送push",
        "name": "send_push_action",
        "params": {
          "project": {
            "label": "项目",
            "field_type": "select",
            "options": [
                {"label": "uangme","value":"uangme"}
            ],
            "type": "",
            "value": "uangme"
          },
           "device_id": {
            "label": "设备id",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          },
          "title": {
            "label": "push标题",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          },
          "content": {
            "label": "push内容",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          },
          "click_action": {
            "label": "跳转路由",
            "field_type": "tag_input",
            "type": "",
            "value": "uangme://api.uangme.com/main"
          }
        }
      },
      {
            "label": "发送站内信",
            "name": "send_news_action",
            "params": {
                "content": {
                    "label": "内容",
                    "field_type": "input",
                    "type": "",
                    "value": ""
                },
                "title": {
                    "label": "标题",
                    "field_type": "input",
                    "type": "",
                    "value": ""
                }
            }
        },
        {"label": "发放优惠券",
            "name": "send_coupon_action",
            "params": {
                "coupon_id": {
                    "label": "优惠券id",
                    "field_type": "input",
                    "type": "int",
                    "value": 0
                }
            }
        },
       {
        "label": "消息触达执行器",
        "name": "strategy_msg_action",
        "params": {
          "project": {
            "label": "项目",
            "field_type": "select",
            "options": [
                {"label": "uangme","value":"uangme"}
            ],
            "type": "",
            "value": "uangme"
          },
          "msg_Key": {
            "label": "场景",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          },
          "device_id": {
            "label": "设备id",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          }
        }
      },
      {
        "label": "uangme发券执行器",
        "name": "uangme_coupon_action",
        "params": {
          "c_id": {
            "label": "券id",
            "field_type": "input",
            "type": "int",
            "value": ""
          },
          "expire_days": {
            "label": "有效期（天）",
            "field_type": "input",
            "type": "int",
            "value": "0"
          },
          "send_rule": {
            "label": "覆盖发卷",
            "field_type": "select",
            "options": [
                {"label": "否","value":0},
                {"label": "是","value":1}
            ],
            "type": "int",
            "value": 0
          },
          "loan_id": {
            "label": "LoanId",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          },
          "plan_id": {
            "label": "PlanId",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          }
        }
      },
       {
        "label": "uangme调额执行器",
        "name": "adjust_amount_action",
        "params": {
          "pid": {
            "label": "产品id",
            "field_type": "select",
            "options": [
                {"label": "分期二期","value":100},
                {"label": "分期三期","value":101},
                {"label": "分期四期","value":102},
                {"label": "分期116","value":116},
                {"label": "分期303","value":303},
                {"label": "分期501","value":501},
                {"label": "payday30天","value":202},
                {"label": "payday21天","value":203},
                {"label": "hartno30天","value":20001}
            ],
            "type": "int",
            "value": ""
          },
          "op": {
            "label": "操作类型",
            "field_type": "select",
            "options": [
                {"label": "提额","value":"add"},
                {"label": "降额","value":"dec"}
            ],
            "type": "",
            "value": ""
          },
          "amount": {
            "label": "调额金额",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          }
        }
      },
      {
        "label": "uangme产品修改执行器",
        "name": "product_modify_action",
        "params": {
          "pid": {
            "label": "产品id",
            "field_type": "select",
            "options": [
                {"label": "分期二期-100","value":100},
                {"label": "分期三期-101","value":101},
                {"label": "分期四期-102","value":102},
                {"label": "分期110","value":110},
                {"label": "分期111","value":111},
                {"label": "分期112","value":112},
                {"label": "分期113","value":113},
                {"label": "分期114","value":114},
                {"label": "分期116","value":116},
                {"label": "分期117","value":117},
                {"label": "分期119","value":119},
                {"label": "分期120","value":120},
                {"label": "分期121","value":121},
                {"label": "分期303","value":303},
                {"label": "分期108","value":108},
                {"label": "分期109","value":109},
                {"label": "payday201","value":201},
                {"label": "payday202","value":202},
                {"label": "payday203","value":203},
                {"label": "payday205","value":205},
                {"label": "payday206","value":206},
                {"label": "payday207","value":207},
                {"label": "payday208","value":208},
                {"label": "payday209","value":209},
                {"label": "payday301","value":301},
                {"label": "payday302","value":302},
                {"label": "payday601","value":601},
                {"label": "hartno30天","value":20001},
                {"label": "horton-10001","value":10001},
                {"label": "horton-10002","value":10002},
                {"label": "ecomm-20001","value":20001},
                {"label": "epp-commerce-10003","value":10003},
                {"label": "epp-commerce-10004","value":10004},
                {"label": "信用卡-701","value":701},
                {"label": "月付-501","value":501},
                {"label": "月付-502","value":502},
                {"label": "月付-503","value":503},
                {"label": "月付-504","value":504},
                {"label": "月付-505","value":505},
                {"label": "月付-506","value":506},
                {"label": "月付-507","value":507},
                {"label": "月付-508","value":508}
            ],
            "type": "int",
            "value": ""
          },
          "is_pass": {
            "label": "产品状态",
            "field_type": "select",
            "options": [
                {"label": "开启","value":"YES"},
                {"label": "关闭","value":"NO"}
            ],
            "type": "",
            "value": ""
          },
          "total_amount": {
            "label": "授信金额",
            "field_type": "tag_input",
            "type": "",
            "value": "0"
          },
          "service_rate": {
            "label": "服务费率",
            "field_type": "input",
            "type": "",
            "value": "0"
          },
          "interest_rate": {
            "label": "利率",
            "field_type": "input",
            "type": "",
            "value": "0"
          },
          "use_default":{
              "label": "默认值填充",
              "field_type": "select",
              "options": [
                {"label": "是","value":true},
                {"label": "否","value":false}
            ],
            "type": "",
            "value": ""
          }
        }
      },
       {
        "label": "渠道转换",
        "name": "appflyer_channel_convert",
        "params": {}
      },
      {
        "label": "老渠道转换",
        "name": "appflyer_channel_old_convert",
        "params": {}
      },
      {
        "label": "刷新渠道",
        "name": "refresh_channel_action",
        "params": {
            "platform_id": {
                "label": "平台id",
                "field_type": "tag_input",
                "type": "",
                "value": 0
            },
            "channel_id": {
                "label": "代理id",
                "field_type": "tag_input",
                "type": "",
                "value": 0
            },
             "os": {
                "label": "设备系统",
                "field_type": "tag_input",
                "type": "",
                "value": 0
            }
        }
      },
      {
            "label": "默认执行器",
            "name": "default_action",
            "params": {
                
            }
      }
    ]
  },
  "lender": {
    "filters": [
      {
        "label": "批次过滤器",
        "name": "task_batch_filter",
        "params": {
          "batch_ids": {
            "label": "批次id(英文,逗号分隔)",
            "field_type": "input",
            "type": "",
            "value": ""
          }
        }
      }
    ],
    "actions": [
      {
        "label": "发送短信",
        "name": "sendmsg_action",
        "params": {
          "sms_channel_type": {
            "label": "短信通道",
            "field_type": "select",
            "options": [
                {"label": "召回营销","value":"6"},
                {"label": "默认","value":"0"},
                {"label": "VIP通道","value":"9"}
            ],
            "type": "",
            "value": "6"
          },
          "sms_content": {
            "label": "短信内容",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          }
        }
      },
      {
            "label": "发送站内信",
            "name": "send_news_action",
            "params": {
                "content": {
                    "label": "内容",
                    "field_type": "input",
                    "type": "",
                    "value": ""
                },
                "title": {
                    "label": "标题",
                    "field_type": "input",
                    "type": "",
                    "value": ""
                }
            }
        },
      {
        "label": "消息触达执行器",
        "name": "strategy_msg_action",
        "params": {
          "project": {
            "label": "项目",
            "field_type": "select",
            "options": [
                {"label": "uangme","value":"uangme"}
            ],
            "type": "",
            "value": "uangme"
          },
          "msg_Key": {
            "label": "场景",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          },
          "device_id": {
            "label": "设备id",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          }
        }
      },
      {
        "label": "lender发券执行器",
        "name": "lender_coupon_action",
        "params": {
          "coupon_type": {
            "label": "优惠券类型",
            "field_type": "select",
            "options": [
                {"label": "体验金","value":1},
                {"label": "现金券","value":2},
                {"label": "加息券","value":3}
            ],
            "type": "",
            "value": 1
          },
          "amount": {
            "label": "金额/利息",
            "field_type": "input",
            "type": "",
            "value": "0"
          },
          "amount_limit": {
            "label": "满多少金额可用",
            "field_type": "input",
            "type": "",
            "value": "0"
          },
          "product_term_limit": {
            "label": "投资满多少天可用",
            "field_type": "input",
            "type": "int",
            "value": "0"
          },
          "product_limit": {
            "label": "投资产品限制\r\n不填无限制\r\n多个产品id用|分隔\r\n如|1|2|3|4|",
            "field_type": "input",
            "type": "",
            "value": ""
          },
          "expire_days": {
            "label": "有效期（天）",
            "field_type": "input",
            "type": "int",
            "value": "0"
          }
        }
      },
     {
        "label": "渠道转换",
        "name": "appflyer_channel_convert",
        "params": {}
      },
      {
        "label": "发送push",
        "name": "send_push_action",
        "params": {
          "project": {
            "label": "项目",
            "field_type": "select",
            "options": [
                {"label": "uangme","value":"uangme"}
            ],
            "type": "",
            "value": "uangme"
          },
           "device_id": {
            "label": "设备id",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          },
          "title": {
            "label": "push标题",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          },
          "content": {
            "label": "push内容",
            "field_type": "tag_input",
            "type": "",
            "value": ""
          },
          "click_action": {
            "label": "跳转路由",
            "field_type": "tag_input",
            "type": "",
            "value": "uangme://api.uangme.com/main"
          }
        }
      }
    ]
  }
}'



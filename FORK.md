# 腾讯广告 SDK fork 差异

本 fork 已同步腾讯官方 `v1.7.87`，提交 `81d6cc04b58d7d5a5ec9ad0137e43b3b6c52fde9`。最近核对日期：2026-09-30。

模型扩展仅保留 `AdgroupsAddRequest.SmartDeliveryPlatform` 与 `SmartDeliverySceneSpec`。2026-01-19 原始中文工单说明保持不变。前者用于 AIM 场景识别；后者为已有 SDK 声明，亿量创建请求已不再使用。

根据腾讯《AIM+ API 使用文档》2026.3 升级说明，小游戏出价统一使用标准 3.0 的 `conversion_id`、`bid_amount` 和 `deep_conversion_worth_rate`。此前自行推定的 30 日专用出价字段已删除，不再保留兼容。30 日变现目标使用转化 ID `10018`，官方对应枚举为 `GOAL_30DAY_MONETIZATION_ROAS`。

回归测试：`go test ./pkg/model/v3 -run 'TestDynamicCreativeBrandsRoundTrip|TestSmartDeliveryDay30RoundTrip' -count=1`。仅验证 JSON，不访问广告接口或数据库。

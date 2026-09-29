# 腾讯广告 SDK fork 差异

本 fork 已同步腾讯官方 `v1.7.87` / `master`，提交 `81d6cc04b58d7d5a5ec9ad0137e43b3b6c52fde9`（2026-08-21），核对日期 2026-09-29。

官方已有的 `CreativeComponents.ChannelsBrand` 和 `SmartDeliveryGoal_APP_REGISTER_30_DAY_MONETIZATION_ROAS` 直接使用官方定义。此前 fork 自行推定的 `SmartDeliveryGoal_DAY30_MONETIZATION` 已删除，不能继续发送对应字符串。

仍保留的模型扩展：

- `AdgroupsAddRequest.SmartDeliveryPlatform` 与 `SmartDeliverySceneSpec`：原有 AIM 创建业务需要，官方最新请求模型仍未包含。2026-01-19 的原始提交记录了腾讯工单回复及补充原因。
- `SmartDeliveryGoalMiniGamePromotionSpec.Day30MonetizationRoi`：`day30_monetization_roi` 为业务按 7 日结构推定的出价字段；官方最新小游戏模型仍未包含，尚无官方资料确认该 JSON 字段与转化 ID `10018` 的完整组合。保留供当前接入使用，离线序列化测试不代表服务端接受。

官方存在“注册 30 日变现”枚举，并不单独证明它适用于 AIM 小游戏 `10018`，也不能据此推断出价字段。未执行真实投放验证。今后若官方增加对应模型，应移除此处重复扩展并同步业务映射。

回归测试：`go test ./pkg/model/v3 -run 'TestDynamicCreativeBrandsRoundTrip|TestSmartDeliveryDay30RoundTrip' -count=1`。这些测试只验证请求 JSON，不访问广告或生产数据库。

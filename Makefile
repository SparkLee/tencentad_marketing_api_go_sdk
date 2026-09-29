ifeq ($(OS),Windows_NT)
GITTAG := v1.0.$(shell powershell.exe -NoProfile -Command "Get-Date -Format yyyyMMddHHmmss")
else
GITTAG := v1.0.$(shell date +%Y%m%d%H%M%S)
endif

.PHONY: gittag
# 创建并推送当前提交的版本标签，与其他广告 SDK 的版本规则保持一致。
gittag:
	git tag $(GITTAG)
	git push origin $(GITTAG)

package main

import "fmt"

type language string

const (
	english language = "en"
	chinese language = "zh-CN"
)

func (l language) normalized() language {
	if l == chinese {
		return chinese
	}
	return english
}

// 用英文原文作为键，让英文界面与翻译共用一处文案引用。
var chineseText = map[string]string{
	"Claude Account Switcher":                       "Claude 账号切换器",
	"Download the new version and restart on it":    "下载新版本并重新启动",
	"Switch the Claude app to this account":         "将 Claude 切换到这个账号",
	"Auto-switch when nearly full":                  "用量接近上限时自动切换",
	"Switch to the account with the most room left": "切换到剩余额度最多的账号",
	"Auto-switch at":                                "自动切换阈值",
	"Custom…":                                       "自定义…",
	"Custom (%.0f%%)…":                              "自定义（%.0f%%）…",
	"Type in your own percentage":                   "输入自定义百分比",
	"Auto-switch to use up expiring weekly":         "自动切换以使用即将重置的每周额度",
	"Switch to an account whose weekly limit resets within a day with half or more of it unused": "切换到每周额度将在一天内重置，且剩余额度不少于一半的账号",
	"Add account…": "添加账号…",
	"Restart Claude on its sign-in screen, your current account is kept": "重新启动 Claude 并打开登录界面，保留当前账号",
	"Import chats from all accounts":                                     "从所有账号导入聊天",
	"Restart Claude so every chat shows up in this account":              "重新启动 Claude，让所有聊天显示在当前账号中",
	"Start at login":                     "开机启动",
	"Open data folder":                   "打开数据文件夹",
	"Quit":                               "退出",
	"Version %s is out, click to update": "发现新版本 %s，点击更新",
	"Open Claude and sign in":            "请打开 Claude 并登录",
	"no accounts yet, sign in to Claude": "还没有账号，请先登录 Claude",
	"usage unknown":                      "用量未知",
	"signed out, add it again":           "登录已失效，请重新添加账号",
	"5h %.0f%%":                          "5 小时 %.0f%%",
	"week %.0f%%":                        "每周 %.0f%%",
	" until %s":                          "，至 %s 重置",
	", resets %s":                        "，%s 重置",
	"the last account":                   "上一个账号",
	"Auto-switched at %s, %s was nearly full":                               "已于 %s 自动切换，%s 的用量接近上限",
	"Auto-switched at %s to use up this account's week before it resets %s": "已于 %s 自动切换，以使用当前账号将在 %s 重置的每周额度",
	"Switching to %s…":                         "正在切换到 %s…",
	"Sign in with the other account in Claude": "请在 Claude 中登录另一个账号",
	"Importing chats…":                         "正在导入聊天…",
	"Switch failed: %v":                        "切换失败：%v",
	"Auto switch failed: %v":                   "自动切换失败：%v",
	"Add account failed: %v":                   "添加账号失败：%v",
	"Import failed: %v":                        "导入失败：%v",
	"Importing chats failed: %v":               "导入聊天失败：%v",
	"Start at login: %v":                       "设置开机启动失败：%v",
	"Could not start Claude: %v":               "无法启动 Claude：%v",
	"unknown account":                          "找不到这个账号",
	"no signed-in account":                     "当前没有已登录的账号",
	"no saved sign-in for %s, use Add account": "%s 没有保存的登录信息，请使用“添加账号”重新登录",
	"Downloading version %s…":                  "正在下载版本 %s…",
	"Update failed: %v":                        "更新失败：%v",
	"Starting version %s…":                     "正在启动版本 %s…",
	"Continue chats?":                          "继续聊天？",
	"Continue your chats?":                     "继续之前的聊天？",
	"These chats were working when Claude closed for the switch. Claude continues the checked ones when it starts again.":  "Claude 为切换账号而关闭时，这些聊天仍在进行。重新启动后，Claude 会继续勾选的聊天。",
	"These chats were working when Claude closed for the switch. Claude continues the selected ones when it starts again.": "Claude 为切换账号而关闭时，这些聊天仍在进行。重新启动后，Claude 会继续选中的聊天。",
	"These chats were working when Claude closed to switch to ":                                                            "Claude 关闭以切换到以下账号时，这些聊天仍在进行：",
	"These chats were working when Claude closed. ":                                                                        "Claude 关闭时，这些聊天仍在进行。",
	". ": "。",
	"Checked chats pick up where they left off.": "勾选的聊天将从中断处继续。",
	"{selected} of {total} selected":             "已选择 {selected} / {total} 个聊天",
	"Select none":                                "取消全选",
	"Select all":                                 "全选",
	"Continue":                                   "继续",
	"Continue 1 chat":                            "继续 1 个聊天",
	"Continue {count} chats":                     "继续 {count} 个聊天",
	"Don't continue":                             "不继续",
	"Continuing in {seconds} s":                  "{seconds} 秒后继续",
	"Waiting for you":                            "等待你的选择",
	"and %d more, which continue too":            "还有 %d 个聊天，也会继续",
	"Continue (%d)":                              "继续（%d）",
	"Claude moves to the account with the most room left once the active one reaches this much of its 5-hour or weekly limit.": "当前账号的 5 小时或每周用量达到此比例时，Claude 将切换到剩余额度最多的账号。",
	"Threshold in percent":               "阈值百分比",
	"Threshold":                          "阈值",
	"Set":                                "设置",
	"Set {value}%":                       "设为 {value}%",
	"Cancel":                             "取消",
	"Enter a whole number from 1 to 99.": "请输入 1 到 99 之间的整数。",
	"It only moves to an account below {room}%, so it does not bounce between full ones.":              "仅切换到用量低于 {room}% 的账号，避免在接近上限的账号之间反复切换。",
	"No account can be {headroom} points below this, so it will not switch.":                           "没有账号的用量能比此阈值低 {headroom} 个百分点，因此不会自动切换。",
	"Auto-switch when the active account reaches this much of its 5-hour or weekly limit, in percent.": "当前账号的 5 小时或每周用量达到以下百分比时，自动切换账号。",
	"Enter a whole number from 1 to 99, for example 85.":                                               "请输入 1 到 99 之间的整数，例如 85。",
}

func (l language) text(source string, args ...any) string {
	if l == chinese {
		if translated, ok := chineseText[source]; ok {
			source = translated
		}
	}
	if len(args) == 0 {
		return source
	}
	return fmt.Sprintf(source, args...)
}

// 状态保留原文和参数，切换语言时可以立即重新生成已有提示。
type localizedMessage struct {
	source string
	args   []any
}

func message(source string, args ...any) localizedMessage {
	return localizedMessage{source: source, args: args}
}

func (m localizedMessage) render(l language) string {
	args := append([]any(nil), m.args...)
	for i, arg := range args {
		if nested, ok := arg.(localizedMessage); ok {
			args[i] = nested.render(l)
		}
	}
	return l.text(m.source, args...)
}

func (m localizedMessage) Error() string {
	return m.render(english)
}

// 弹窗只传所选语言的翻译，英文直接使用页面中的原文。
func (l language) translations() map[string]string {
	if l == chinese {
		return chineseText
	}
	return nil
}

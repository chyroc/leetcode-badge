# Leetcode Badge

[![Build Status](https://travis-ci.org/Chyroc/leetcode-badge.svg?branch=master)](https://travis-ci.org/Chyroc/leetcode-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/Chyroc/leetcode-badge)](https://goreportcard.com/report/github.com/Chyroc/leetcode-badge)
[![codecov](https://codecov.io/gh/Chyroc/leetcode-badge/branch/master/graph/badge.svg)](https://codecov.io/gh/Chyroc/leetcode-badge)

[![leetcode badge](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Ranking-{{.ranking}}-green.svg&refresh=true)](https://github.com/Chyroc/leetcode-badge)
[![leetcode badge](https://leetcode-badge.chyroc.cn/?name=chyroc&refresh=true)](https://github.com/Chyroc/leetcode-badge)
[![leetcode badge](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Question-{{.solved_question_rate}}-{{%20if%20le%20.solved_question_rate_float%200.3}}red{{%20else%20if%20le%20.solved_question_rate_float%200.6}}yellow{{%20else%20}}green{{%20end%20}}.svg&refresh=true)](https://github.com/Chyroc/leetcode-badge)
[![leetcode badge](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Submission-{{.accepted_submission_rate}}-{{%20if%20le%20.accepted_submission_rate_float%200.3}}red{{%20else%20if%20le%20.solved_question_rate_float%200.6}}yellow{{%20else%20}}green{{%20end%20}}.svg&refresh=true)](https://github.com/Chyroc/leetcode-badge)

[English Document](./README_en.md)

Leetcode Badge是一个展示leetcode通过情况徽标的项目。

svg绘制依赖于[shields.io](http://shields.io/)，所以任何shields.io支持的语法，这里都适用。

## 示例

* 默认风格 ![](https://leetcode-badge.chyroc.cn/?name=chyroc&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc`

  > 注意：这里的颜色是会变化的 红（低于等于30％），黄（低于等于60％），绿（其他）

* 排名 ![leetcode badge](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Ranking-{{.ranking}}-green.svg&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Ranking-{{.ranking}}-green.svg`

* 通过题目/总题目数 ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-green.svg&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-green.svg`

* 通过的提交/总提交数 ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Accepted/Total-{{.accepted_submission}}/{{.all_submission}}-green.svg&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Accepted/Total-{{.accepted_submission}}/{{.all_submission}}-green.svg`

* 通过题目/总题目数 + 自定义的style ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-green.svg?style=flat-square&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-green.svg?style=flat-square`

* 通过题目/总题目数 + 自定义的颜色 ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-red.svg&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-red.svg`

* 计算不同的比例显示不同的颜色 ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Solved/Total-{{.solved_question}}/{{.all_question}}-{{if%20le%20.solved_question_rate_float%200.3}}red.svg{{else%20if%20le%20.solved_question_rate_float%200.6}}yellow.svg{{else}}green.svg{{end}}&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode | Solved/Total-{{.solved_question}}/{{.all_question}}-{{ if le .solved_question_rate_float 0.3}}red.svg{{ else if le .solved_question_rate_float 0.6}}yellow.svg{{ else }}green.svg{{ end }}`

  > 注意：这里的颜色是会变化的 红（低于等于30％），黄（低于等于60％），绿（其他）

* 通过题目/总题目 比例 ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Question-{{.solved_question_rate}}-{{%20if%20le%20.solved_question_rate_float%200.3}}red{{%20else%20if%20le%20.solved_question_rate_float%200.6}}yellow{{%20else%20}}green{{%20end%20}}.svg&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode | Question-{{.solved_question_rate}}-{{ if le .solved_question_rate_float 0.3}}red{{ else if le .solved_question_rate_float 0.6}}yellow{{ else }}green{{ end }}.svg`

* 通过的提交/总提交数 比例 ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Submission-{{.accepted_submission_rate}}-{{%20if%20le%20.accepted_submission_rate_float%200.3}}red{{%20else%20if%20le%20.solved_question_rate_float%200.6}}yellow{{%20else%20}}green{{%20end%20}}.svg&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode | Submission-{{.accepted_submission_rate}}-{{ if le .accepted_submission_rate_float 0.3}}red{{ else if le .solved_question_rate_float 0.6}}yellow{{ else }}green{{ end }}.svg`



## 语法

### 使用querystring传递参数
* name：leetcode用户名
* leetcode_badge_style：自定义的badge显示格式

### 使用go的模板作为leetcode_badge_style语法

#### go的模板语法

* 参考：https://godoc.org/text/template
* 简述
  * `{{ .xxx }}`可以引用下面的6个变量
  * `{{ le .xx 0.3 }} a {{ else if le 0.6 }} b {{ else }} c` xx小于等于0.3返回a，小于等于0.6返回b，否则返回c

#### 可以使用go的模板语法使用6个变量：
* {{.ranking}} 排名（整数）
* {{.accepted_submission}} 通过的提交的个数（整数）
* {{.all_submission}} 所有的提交的个数（整数）
* {{.solved_question}} 通过的题目的个数（整数）
* {{.all_question}} 所有的题目的个数（整数）
* {{.solved_question_rate_float}} 通过的题目占总题目数的比例（小数）
* {{.accepted_submission_rate_float}} 提交通过的占总提交数的比例（小数）
* {{.solved_question_rate}} 通过的题目占总题目数的比例（形如23％的字符串）
* {{.accepted_submission_rate}} 提交通过的占总提交数的比例（形如23％的字符串）

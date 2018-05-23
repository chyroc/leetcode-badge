# Leetcode Badge

[英文文档](./README_en.md)

Leetcode Badge是一个展示leetcode badge的项目。

svg绘制依赖于[shields.io](shields.io)，所以任何shields.io支持的语法，这里都使用

## 使用说明

## 默认形式

https://leetcode-badge.chyroc.cn/?name=chyroc

![Leetcode Badge](https://leetcode-badge.chyroc.cn/?name=chyroc)


## 自定义

### 自定义语法

可以使用go的模板语法使用6个变量：
* {{.accepted_submission}} 通过的提交的个数
* {{.all_submission}} 所有的提交的个数
* {{.solved_question}} 通过的题目的个数
* {{.all_question}} 所有的题目的个数
* {{.solved_question_rate}} 通过的题目占总题目数的比例
* {{.accepted_submission_rate}} 提交通过的占总提交数的比例

然后将shields.io支持的语法放在query string leetcode_badge_style 后


### 示例

* 通过题目/总题目数
  * `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-green.svg`
  * ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-green.svg)

* 通过题目/总题目数 + 自定义的style
  * `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-green.svg?style=flat-square`
  * ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-green.svg?style=flat-square)

* 通过题目/总题目数 + 自定义的颜色
  * `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-red.svg`
  * ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-red.svg)

*通过题目/总题目 比例
 * `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=SolvedRate-{{.solved_question_rate}}-green.svg`
 * ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=SolvedRate-{{.solved_question_rate}}-green.svg)

* 通过的提交/总提交数
  * `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Accepted/Total-{{.accepted_submission}}/{{.all_submission}}-green.svg`
  * ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Accepted/Total-{{.accepted_submission}}/{{.all_submission}}-green.svg)


# Leetcode Badge

[![leetcode badge](https://leetcode-badge.chyroc.cn/?name=chyroc&refresh=true)](https://github.com/Chyroc/leetcode-badge)
[![leetcode badge](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Question-{{.solved_question_rate}}-{{%20if%20le%20.solved_question_rate_float%200.3}}red{{%20else%20if%20le%20.solved_question_rate_float%200.6}}yellow{{%20else%20}}green{{%20end%20}}.svg&refresh=true)](https://github.com/Chyroc/leetcode-badge)
[![leetcode badge](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Submission-{{.accepted_submission_rate}}-{{%20if%20le%20.accepted_submission_rate_float%200.3}}red{{%20else%20if%20le%20.solved_question_rate_float%200.6}}yellow{{%20else%20}}green{{%20end%20}}.svg&refresh=true)](https://github.com/Chyroc/leetcode-badge)

[中文文档](./README.md)

Leetcode Badge is a project that showcases the leetcode pass situation logo.

Svg drawing relies on [shields.io](http://shields.io/), so any shields.io supported syntax is applicable here.

## Example

* Default style ![](https://leetcode-badge.chyroc.cn/?name=chyroc&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc`

  > Note: The color here is red (less than or equal to 30%), yellow (less than or equal to 60%), green (other)

* Pass the title/total number of topics ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{solve_question}}/{{all_question}}-green.svg&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-green.svg`

* Number of submitted/total submissions ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Accepted/Total-{{accepted_submission}}/{{all_submission}}-green.svg&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Accepted/Total-{{accepted_submission}}/{{.all_submission}}-green.svg`

* Pass the title/total number of topics + custom style ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-green.svg?style=flat-square&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-green.svg?style=flat-square`

* Pass the title/total number of topics + custom colors ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-red.svg&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Solved/Total-{{.solved_question}}/{{.all_question}}-red.svg`

* Calculate different proportions to show different colors ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Solved/Total-{{solve_question}}/{{.all_question}}-{{if%20le%20.solved_question_rate_float%200.3}}red.svg{{else%20if%20le%20.solved_question_rate_float%200.6}}yellow.svg{{else}}green.svg{{end}}&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode | Solved/Total-{{.solved_question}}/{{.all_question}}-{{ if le .solved_question_rate_float 0.3}}red. Svg{{ else if le .solved_question_rate_float 0.6}}yellow.svg{{ else }}green.svg{{ end }}`

  > Note: The color here is red (less than or equal to 30%), yellow (less than or equal to 60%), green (other)

* Pass the title/total title ratio ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Question-{{.solved_question_rate}}-{{%20if%20le%20.solved_question_rate_float%200.3}}red{{%20else%20if%20le%20.solved_question_rate_float%200.6}}yellow{{%20else%20}}green{{%20end%20}}.svg&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode | Question-{{.solved_question_rate}}-{{ if le .solved_question_rate_float 0.3}}red{{ else if le .solved_question_rate_float 0.6}} Yellow{{ else }}green{{ end }}.svg`

* Proportion of submitted/total submissions ![](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Submission-{{.accepted_submission_rate}}-{{%20if%20le%20.accepted_submission_rate_float%200.3}}red{{%20else%20if%20le%20.solved_question_rate_float%200.6}}yellow{{%20else%20}}green{{%20end%20}}.svg&refresh=true)
  > `https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode | Submission-{{.accepted_submission_rate}}-{{ if le .accepted_submission_rate_float 0.3}}red{{ else if le .solved_question_rate_float 0.6}} Yellow{{ else }}green{{ end }}.svg`

## Syntax

### Using querystring to pass parameters
* name: leetcode username
* leetcode_badge_style: custom badge display format

### Using go's template as the leetcode_badge_style syntax

#### go's template syntax

* Reference: https://godoc.org/text/template
* Brief description
  * `{{ .xxx }}` can reference the following 6 variables
  * `{{ le .xx 0.3 }} a {{ else if le 0.6 }} b {{else }} c` xx is less than or equal to 0.3 returns a, less than or equal to 0.6 returns b, otherwise returns c

#### You can use go's template syntax to use 6 variables:
* {{.accepted_submission}} Number of submitted submissions (integer)
* {{.all_submission}} The number of all submissions (integer)
* {{.solved_question}} Number of questions passed (integer)
* {{.all_question}} The number of all questions (integer)
* {{.solved_question_rate_float}} Proportion of the total number of questions passed by the topic (decimal)
* {{.accepted_submission_rate_float}} Proportion of total submissions submitted (decimal)
* {{.solved_question_rate}} The percentage of questions passed  (as a string of 23%)
* {{.accepted_submission_rate}} Proportion of submissions passed (as a string of 23%)
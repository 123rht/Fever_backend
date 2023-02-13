# Fever_backend
该项目主要是对发热人员进行统计，主要范围是市，管理员，院长，医生的工作责任不同所以权限也不同，三个角色都可使用，但是对应展示的内容和权限不同
这个项目使用gin框架开发的前后端分离项目，这是后端部分。
使用了gorm,casbin等技术
自己实现了一套jwt-token鉴权
casbin权限管理，三个不同角色分别是管理员，院长，医生，三个角色对应的权限不同
个人负责的是院长的模块，此模块与医生关联，所以在进行相应的crud时需要对医生模块进行改动

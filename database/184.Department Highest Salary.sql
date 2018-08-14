# Write your MySQL query statement below
select b.Name as Department, a.Name as Employee,  a.Salary
from Employee a 
inner join (select DepartmentId,max(Salary) as Salary from Employee group by DepartmentId)c
on a.DepartmentId=c.DepartmentId and a.Salary = c.Salary
inner join Department b
on a.DepartmentId = b.id
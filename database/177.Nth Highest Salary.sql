CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      # Write your MySQL query statement below.
      
      select Salary from (
      select Salary,case when @last=Salary then @rank else @rank:=@rank+1 end as Rank, @last:=Salary
      from Employee a,(select @rank:=0, @last:=null)b
      order by Salary desc)t
      where Rank=N order by Salary asc limit 1   
  );
END
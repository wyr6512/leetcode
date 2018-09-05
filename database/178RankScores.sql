# Write your MySQL query statement below
select Score, Rank 
from (
select a.Score, convert(case when @lastscore = a.Score then @rank else @rank := @rank+1 end, SIGNED) as Rank,@lastscore := a.Score
from Scores a,(select @rank := 0, @lastscore := null)b
order by Score desc
)t

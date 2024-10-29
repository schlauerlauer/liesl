-- name: GetNodes :many
select *
from nodes;

-- name: GetNode :one
select
    ID
from nodes
where id = ?
limit 1;

-- name: GetEdges :many
select
    target
from edges
where source = ?;

-- name: InsertNode :exec
insert into nodes (ID)
values (?)
on conflict (ID) do nothing;

-- name: InsertEdge :exec
insert into edges (source, target)
values (?, ?);

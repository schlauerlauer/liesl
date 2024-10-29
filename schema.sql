create table if not exists nodes (
  id   text primary key
);

create index if not exists nodes_id on nodes(id);

create table if not exists edges (
  source text not null,
  target text not null,
  unique(source,target) on conflict replace,
  foreign key(source) references nodes(id),
  foreign key(target) references nodes(id)
  --properties?
);

create index if not exists edges_source on edges(source);
create index if not exists edges_target on edges(target);

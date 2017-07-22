insert into lists(name, list_type, created_at) values("groceries", "shoppingList", TIMESTAMP(curdate(), curtime()));
insert into categories(name, list_id, created_at) values("Produce", 1, TIMESTAMP(curdate(), curtime()));
insert into categories(name, list_id, created_at) values("Dairy", 1, TIMESTAMP(curdate(), curtime()));

insert into items(name, quantity, measure, status, list_id, category_id, created_at)
values("apples", 5, null, "needed", 1, 1, TIMESTAMP(curdate(), curtime()));

insert into items(name, quantity, measure, status, list_id, category_id, created_at)
values("bananas", 1, "bunch", "needed", 1, 1, TIMESTAMP(curdate(), curtime()));

insert into items(name, quantity, measure, status, list_id, category_id, created_at)
values("Soy milk", 1, null, "needed", 1, 2, TIMESTAMP(curdate(), curtime()));

select * from lists;
select * from categories;
select * from items;


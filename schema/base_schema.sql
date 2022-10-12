create table if not exists available_virtual_bond (
	virtual_bond_id varchar(255) not null unique,
    virtual_bond int not null unique,
    primary key (virtual_bond_id, virtual_bond)
);
create table if not exists virtual_bond (
	uid varchar(255) not null unique,
	virtual_bond_id varchar(255),
	a_faction_id varchar(255) not null,
    b_faction_id varchar(255) not null,
    end_bond_time varchar(100) not null,
    primary key (uid),
    foreign key (virtual_bond_id) references available_virtual_bond(virtual_bond_id)
);





go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/xorwise/effective-mobile-tz/bootstrap"
)

func main() {
	env := bootstrap.NewEnv()
	conn, err := bootstrap.NewDatabaseConnection(env)
	if err != nil {
		log.Fatal(err)
	}

	err = bootstrap.MigrateDatabase(conn)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(context.Background(), `
		INSERT INTO cars (reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic)
		VALUES
			('A123AA123', 'Ford', 'Focus', 2020, 'John', 'Doe', 'Smith'),
			('B456BB456', 'Toyota', 'Camry', 2018, 'Jane', 'Doe', 'Smith'),
			('C789CC789', 'Nissan', 'Altima', 2021, 'Bob', 'Smith', 'Johnson'),
			('D012DD012', 'BMW', 'M5', 2019, 'Alice', 'White', 'Brown'),
			('E345EE345', 'Hyundai', 'Elantra', 2022, 'Charlie', 'Brown', 'Davis'),
			('F678FF678', 'Chevrolet', 'Cruze', 2017, 'David', 'Davis', 'Wilson'),
			('G901GG901', 'Kia', 'Soul', 2016, 'Emily', 'Wilson', 'Taylor'),
			('H234HH234', 'BMW', 'X5', 2023, 'Frank', 'Taylor', 'Thomas'),
			('I567II567', 'Mercedes-Benz', 'C-Class', 2022, 'Grace', 'Thomas', 'Jackson'),
			('J890JJ890', 'Audi', 'A4', 2014, 'Henry', 'Jackson', 'White'),
			('K123KK123', 'Porsche', 'Cayenne', 2013, 'Ivy', 'Roberts', 'Harris'),
			('L456LL456', 'Mazda', 'CX5', 2012, 'Jack', 'Harris', 'Martin'),
			('M789MM789', 'BMW', 'I8', 2011, 'Kate', 'Martin', 'Thompson'),
			('N012NN012', 'Lamborghini', 'Aventador', 2022, 'Liam', 'Thompson', 'Garcia'),
			('O345OO345', 'Bugatti', 'Chiron', 2006, 'Mia', 'Garcia', 'Martinez'),
			('P678PP678', 'Dodge', 'Challenger', 2008, 'Natalie', 'Martinez', 'Roberts'),
			('Q901QQ901', 'Jaguar', 'XJ', 2007, 'Olivia', 'Hernandez', 'Lewis'),
			('R234RR234', 'Tesla', 'Model S', 2006, 'Penelope', 'Lewis', 'Walker'),
			('S567SS567', 'BMW', 'X6', 2005, 'Quinn', 'Walker', 'Hernandez'),
			('T890TT890', 'Ferrari', '488 GTB', 2004, 'Rachel', 'Hernandez', 'Lopez'),
			('U123UU123', 'Lamborghini', 'Aventador SVJ', 2006, 'Samantha', 'Lopez', 'Perez'),
			('V456VV456', 'Bugatti', 'Chiron Super Sport', 2002, 'Sofia', 'Perez', 'Robinson'),
			('W789WW789', 'Dodge', 'Challenger SRT', 2001, 'Tobias', 'Robinson', 'Walker'),
			('X012XX012', 'Jaguar', 'XJ12', 2000, 'Ursula', 'Walker', 'Harris');
	`)
	if err != nil {
		log.Fatal(err)
	}
	conn.Close(context.Background())
	fmt.Println("Done")
}

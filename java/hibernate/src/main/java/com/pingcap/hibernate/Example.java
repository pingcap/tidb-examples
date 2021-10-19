/* 
 * Hibernate, Relational Persistence for Idiomatic Java
 * 
 * JBoss, Home of Professional Open Source
 * Copyright 2014 Red Hat Inc. and/or its affiliates and other contributors
 * as indicated by the @authors tag. All rights reserved.
 * See the copyright.txt in the distribution for a
 * full listing of individual contributors.
 *
 * This copyrighted material is made available to anyone wishing to use,
 * modify, copy, or redistribute it subject to the terms and conditions
 * of the GNU Lesser General Public License, v. 2.1.
 * This program is distributed in the hope that it will be useful, but WITHOUT A
 * WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
 * PARTICULAR PURPOSE.  See the GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License,
 * v.2.1 along with this distribution; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
 * MA  02110-1301, USA.
 */
package com.pingcap.hibernate;

import org.hibernate.Session;
import org.hibernate.SessionFactory;
import org.hibernate.boot.registry.StandardServiceRegistryBuilder;
import org.hibernate.cfg.Configuration;
import org.hibernate.query.Query;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import java.util.Iterator;
import java.util.List;

public class Example {

	public static enum Gender {
		MALE, FEMALE
	}

	/**
	 * The Class User.
	 */
	@Entity
	public static class User {
		/** The id. */
		@Id
		@GeneratedValue(strategy=GenerationType.IDENTITY)
		private int userId;

		private String name;

		private Gender gender;

		public User() {}

		public User(String name, Gender gender) {
			setName(name);
			setGender(gender);
		}

		public int getUserId() {
			return userId;
		}

		public String getName() {
			return name;
		}

		public void setName(String name) {
			this.name = name;
		}

		public Gender getGender() {
			return gender;
		}

		public void setGender(Gender gender) {
			this.gender = gender;
		}
	}

	@Entity
	public static class Order {
		/** The order id. */
		@Id
		@GeneratedValue(strategy=GenerationType.IDENTITY)
		private int orderId;

		/** The user id related to this order. */
		private int userId;

		/** The price */
		private double price;

		public Order() {}

		public Order(int userId, double price) {
			setUserId(userId);
			setPrice(price);
		}

		public int getOrderId() {
			return orderId;
		}

		public int getUserId() {
			return userId;
		}

		public void setUserId(int userId) {
			this.userId = userId;
		}

		public double getPrice() {
			return price;
		}

		public void setPrice(double price) {
			this.price = price;
		}
	}


	/**
	 * The main method.
	 *
	 * @param args the arguments
	 */
	public static void main(String[] args) {
		Session session = openSession();
		try {
			User tom = new User("Tom", Gender.MALE);
			persist(session, tom);

			User jack = new User("Jack", Gender.MALE);
			persist(session, jack);

			Order order1 = new Order(tom.getUserId(), 100.0);
			Order order2 = new Order(tom.getUserId(), 200.0);
			Order order3 = new Order(jack.getUserId(), 300.0);

			persist(session, order1);
			persist(session, order2);
			persist(session, order3);

			order1.setPrice(500);
			persist(session, order1);

			delete(session, order3);

			Query query = session.createQuery("SELECT u.name FROM Example$User u JOIN Example$Order o ON u.userId = o.userId GROUP BY u.name HAVING SUM(o.price) > 500");
			List list = query.list();
			Iterator it = list.iterator();
			while (it.hasNext()) {
				System.out.println("User name: " + (String) it.next());
			}
		} finally {
			session.close();
		}
	}
	
	/**
	 * Insert user.
	 *
	 * @param obj the user
	 * @throws Exception the exception
	 */
	private static void persist(Session session, Object obj) {
		session.getTransaction().begin();
		session.persist(obj);
		session.getTransaction().commit();
	}

	/**
	 * Insert user.
	 *
	 * @param obj the user
	 * @throws Exception the exception
	 */
	private static void delete(Session session, Object obj) {
		session.getTransaction().begin();
		session.delete(obj);
		session.getTransaction().commit();
	}

	/** The session factory. */
	private static SessionFactory sessionFactory = null;

	/**
	 * Open session.
	 *
	 * @return the session
	 */
	private static Session openSession() {
		if (sessionFactory == null) {
			final Configuration configuration = new Configuration();
			configuration.addAnnotatedClass( User.class );
			configuration.addAnnotatedClass( Order.class );

			sessionFactory = configuration.buildSessionFactory( new StandardServiceRegistryBuilder().build() );
		}
		return sessionFactory.openSession();
	}
}

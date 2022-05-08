package com.bookstore.ecommerce.repository;

import javax.persistence.Persistence;
import lombok.Getter;

public class EntityManagerFactory {
  private javax.persistence.EntityManagerFactory factory;

  public EntityManagerFactory(javax.persistence.EntityManagerFactory factory) {
    this.factory = factory;
  }

  public EntityManager create() {
    return new EntityManager(this.factory.createEntityManager());
  }

  public class EntityManager implements AutoCloseable {
    @Getter
    private javax.persistence.EntityManager manager;

    public EntityManager(javax.persistence.EntityManager manager) {
      this.manager = manager;
    }

    @Override
    public void close() {
      this.manager.close();
    }
  }

  private static EntityManagerFactory onlyOneFactory;

  public static EntityManagerFactory getInstance() {
    if (onlyOneFactory == null) {
      onlyOneFactory =
        new EntityManagerFactory(Persistence.createEntityManagerFactory("cassandra_pu"));
    }

    return onlyOneFactory;
  }
}

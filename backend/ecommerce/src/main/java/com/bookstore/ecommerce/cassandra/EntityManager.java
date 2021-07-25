package com.bookstore.ecommerce.cassandra;

import javax.persistence.Persistence;

import lombok.Getter;

public class EntityManager implements AutoCloseable {
  @Getter
  private javax.persistence.EntityManager manager;

  public EntityManager() {
    this.manager = Persistence.createEntityManagerFactory("cassandra_pu").createEntityManager();
  }

  @Override
  public void close() {
    this.manager.close();
  }
}

package com.bookstore.ecommerce.repository.cassandra;

import javax.persistence.EntityTransaction;
import javax.persistence.FlushModeType;

import com.bookstore.ecommerce.app.repository.Transaction;

import lombok.Getter;

public class CassandraTransaction implements Transaction {
  @Getter
  private EntityManager manager;
  private EntityTransaction tx;

  public CassandraTransaction(EntityManager manager) {
    this.manager = manager;
    this.tx = manager.getManager().getTransaction();
  }

  public void begin() {
    this.manager.getManager().setFlushMode(FlushModeType.COMMIT);
    this.tx.begin();
  }

  @Override
  public void commit() throws Exception {
    this.tx.commit();
  }

  @Override
  public void rollback() throws Exception {
    this.tx.rollback();
  }
}

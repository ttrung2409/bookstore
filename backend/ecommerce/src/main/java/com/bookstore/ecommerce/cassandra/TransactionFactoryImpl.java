package com.bookstore.ecommerce.cassandra;

import java.util.concurrent.Callable;
import java.util.concurrent.CompletableFuture;

import org.springframework.stereotype.Component;

import com.bookstore.ecommerce.app.repository.TransactionFactory;
import com.ea.async.Async;

@Component
public class TransactionFactoryImpl implements TransactionFactory {
  @Override
  public <T> T runInTransaction(Callable<CompletableFuture<T>> func) throws Exception {
    try (var manager = new EntityManager()) {
      var transaction = manager.getManager().getTransaction();

      try {
        transaction.begin();
        var result = Async.await(func.call());
        transaction.commit();

        return result;
      } catch (Exception e) {
        transaction.rollback();
        throw e;
      }
    }
  }
}

package com.bookstore.ecommerce.repository.cassandra;

import java.util.concurrent.CompletableFuture;
import org.springframework.stereotype.Component;
import com.bookstore.ecommerce.app.repository.TransactionFactory;
import com.bookstore.ecommerce.app.repository.TransactionalFunc;

@Component
public class CassandraTransactionFactory implements TransactionFactory {
  @Override
  public <R> CompletableFuture<R> runInTransaction(TransactionalFunc<CompletableFuture<R>> func)
    throws Exception {
    try (var manager = new EntityManager()) {
      var transaction = new CassandraTransaction(manager);

      try {
        transaction.begin();
        var result = func.apply(transaction).join();
        transaction.commit();

        return CompletableFuture.completedFuture(result);
      } catch (Exception e) {
        transaction.rollback();
        throw e;
      }
    }
  }
}

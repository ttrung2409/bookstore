package com.bookstore.ecommerce.repository;

import java.util.concurrent.CompletableFuture;
import org.springframework.stereotype.Component;
import lombok.var;
import com.bookstore.ecommerce.app.repository.TransactionFactory;
import com.bookstore.ecommerce.app.repository.TransactionalFunc;

@Component
public class CassandraTransactionFactory implements TransactionFactory {
  @Override
  public <R> CompletableFuture<R> runInTransaction(TransactionalFunc<CompletableFuture<R>> func)
    throws Exception {
    var transaction =
      new CassandraTransaction(EntityManagerFactory.getInstance().create());

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

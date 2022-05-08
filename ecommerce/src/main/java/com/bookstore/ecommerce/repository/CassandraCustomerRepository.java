package com.bookstore.ecommerce.repository;

import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.domain.Customer;
import com.bookstore.ecommerce.app.repository.CustomerRepository;
import com.bookstore.ecommerce.app.repository.Transaction;
import org.springframework.stereotype.Component;
import lombok.var;

@Component
public class CassandraCustomerRepository implements CustomerRepository {

  @Override
  public CompletableFuture<Void> createIfNotExist(Customer customer, Transaction tx)
    throws Exception {
    try (var manager = EntityManagerFactory.getInstance().create()) {
      var existingCustomer = manager.getManager()
        .find(com.bookstore.ecommerce.app.domain.Customer.class, customer.getState().getPhone());

      if (existingCustomer != null) {
        return null;
      }

      manager.getManager().persist(customer.getState());

      return null;
    }
  }
}

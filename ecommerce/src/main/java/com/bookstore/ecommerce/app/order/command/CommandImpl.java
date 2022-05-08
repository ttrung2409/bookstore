package com.bookstore.ecommerce.app.order.command;

import java.util.ArrayList;
import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.repository.CustomerRepository;
import com.bookstore.ecommerce.app.repository.OrderRepository;
import com.bookstore.ecommerce.app.repository.TransactionFactory;
import org.springframework.stereotype.Component;
import lombok.var;

@Component
public class CommandImpl implements Command {
  private final TransactionFactory transactionFactory;
  private final OrderRepository orderRepository;
  private final CustomerRepository customerRepository;

  public CommandImpl(
    TransactionFactory transactionFactory,
    OrderRepository orderRepository,
    CustomerRepository customerRepository) {
    this.transactionFactory = transactionFactory;
    this.orderRepository = orderRepository;
    this.customerRepository = customerRepository;
  }

  @Override
  public CompletableFuture<String> createOrder(CreateOrderRequest request) throws Exception {
    final var books = new ArrayList<com.bookstore.ecommerce.app.domain.data.Book>();
    for (final var book : request.getBooks()) {
      books.add(book.toDataObject());
    }

    var customer =
      new com.bookstore.ecommerce.app.domain.Customer(request.getCustomer().toDataObject());

    var order = new com.bookstore.ecommerce.app.domain.Order(customer.getState(), books);

    return this.transactionFactory.runInTransaction(tx -> {
      this.customerRepository.createIfNotExist(customer, tx);

      this.orderRepository.create(order, tx).get();

      return CompletableFuture.completedFuture(order.getState().getId());
    });
  }

  @Override
  public CompletableFuture<Void> cancelOrder(String orderId) throws Exception {
    return this.transactionFactory.runInTransaction(tx -> {
      var order = this.orderRepository.get(orderId, tx).get();
      order.cancel();
      this.orderRepository.update(order, tx).get();

      return CompletableFuture.completedFuture(null);
    });
  }
}

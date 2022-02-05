package com.bookstore.ecommerce.app.order.command;

import java.util.ArrayList;
import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.order.command.dto.CreateOrderRequest;
import com.bookstore.ecommerce.app.repository.OrderRepository;
import com.bookstore.ecommerce.app.repository.TransactionFactory;
import org.springframework.stereotype.Component;
import lombok.var;

@Component
public class CommandImpl implements Command {
  private TransactionFactory transactionFactory;
  private OrderRepository orderRepository;

  public CommandImpl(TransactionFactory transactionFactory,
      OrderRepository orderRepository) {
    this.transactionFactory = transactionFactory;
    this.orderRepository = orderRepository;
  }

  @Override
  public CompletableFuture<String> createOrder(CreateOrderRequest request) throws Exception {
    final var books = new ArrayList<com.bookstore.ecommerce.app.domain.data.Book>();
    for (final var book : request.getBooks()) {
      books.add(book.toDataObject());
    }

    var order = new com.bookstore.ecommerce.app.domain.Order(
        request.getCustomer().toDataObject(),
        books);

    return this.transactionFactory.runInTransaction(tx -> {
      this.orderRepository.create(order, tx).join();

      return CompletableFuture.completedFuture(order.getState().getId());
    });
  }
}

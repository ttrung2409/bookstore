package com.bookstore.ecommerce.app.order.command;

import java.util.ArrayList;

import com.bookstore.ecommerce.app.repository.OrderRepository;
import com.bookstore.ecommerce.app.repository.TransactionFactory;

import org.springframework.stereotype.Component;

import lombok.experimental.ExtensionMethod;

@Component
@ExtensionMethod({ BookMapper.class, CustomerMapper.class })
public class CreateOrderCommandImpl implements CreateOrderCommand {
  private TransactionFactory transactionFactory;
  private OrderRepository orderRepository;

  public CreateOrderCommandImpl(TransactionFactory transactionFactory, OrderRepository orderRepository) {
    this.transactionFactory = transactionFactory;
    this.orderRepository = orderRepository;
  }

  @Override
  public void execute(CreateOrderRequest request) throws Exception {
    final var books = new ArrayList<com.bookstore.ecommerce.app.domain.data.Book>();
    for (final var book : request.getBooks()) {
      books.add(book.toDataObject());
    }

    final var order = new com.bookstore.ecommerce.app.domain.Order(request.getCustomer().toDataObject(),
        books.toArray(new com.bookstore.ecommerce.app.domain.data.Book[books.size()]));

    this.transactionFactory.runInTransaction(tx -> {
      return this.orderRepository.create(order.getState(), tx);
    });
  }
}

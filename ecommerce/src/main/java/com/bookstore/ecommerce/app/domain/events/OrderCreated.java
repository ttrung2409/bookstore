package com.bookstore.ecommerce.app.domain.events;

import com.bookstore.ecommerce.app.domain.Event;
import com.bookstore.ecommerce.app.domain.data.Order;
import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
public class OrderCreated extends Event {
  @Getter
  Order order;
}

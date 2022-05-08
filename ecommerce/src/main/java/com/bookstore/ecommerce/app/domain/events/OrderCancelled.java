package com.bookstore.ecommerce.app.domain.events;

import com.bookstore.ecommerce.app.domain.Event;
import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
public class OrderCancelled extends Event {
  @Getter
  String orderId;
}

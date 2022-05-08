package com.bookstore.ecommerce.app.domain;

import java.util.UUID;
import com.google.common.base.Strings;
import lombok.Getter;
import lombok.var;

public class Customer {
  @Getter
  private com.bookstore.ecommerce.app.domain.data.Customer customer;

  public Customer(com.bookstore.ecommerce.app.domain.data.Customer customer) {
    var cloned = customer.clone();
    if (Strings.isNullOrEmpty(cloned.getId())) {
      cloned.setId(UUID.randomUUID().toString());
    }

    this.customer = cloned;
  }

  public com.bookstore.ecommerce.app.domain.data.Customer getState() {
    return this.customer.clone();
  }
}

package com.bookstore.ecommerce.cassandra;

import com.datastax.oss.driver.api.mapper.annotations.PartitionKey;

public class Order extends com.bookstore.ecommerce.app.data.Order {
  @Override
  @PartitionKey public String getId() {
    return super.getId();
  }  
}

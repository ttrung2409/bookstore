package com.bookstore.ecommerce.app.domain;

import java.util.ArrayList;
import java.util.List;

public class EventSource {
  protected List<Event> pendingEvents = new ArrayList<Event>();
}

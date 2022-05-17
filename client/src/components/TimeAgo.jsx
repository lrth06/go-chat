export default function TimeAgo({ date }) {
  const incoming = new Date(date);
  const now = new Date();
  let diff = now - incoming;
  const years = Math.floor(diff / (1000 * 60 * 60 * 24 * 365));
  const months = Math.floor(diff / (1000 * 60 * 60 * 24 * 30));
  const days = Math.floor(diff / (1000 * 60 * 60 * 24));
  const hours = Math.floor(diff / (1000 * 60 * 60));
  const minutes = Math.floor(diff / (1000 * 60));
  const seconds = Math.floor(diff / 1000);
  if (years > 0) {
    if (years === 1) {
      return '1 year ago';
    }
    return years + ' years ago';
  }
  if (months > 0) {
    if (months === 1) {
      return '1 month ago';
    }
    return months + ' months ago';
  }
  if (days > 0) {
    if (days === 1) {
      return '1 day ago';
    }
    return days + ' days ago';
  }
  if (hours > 0) {
    if (hours === 1) {
      return '1 hour ago';
    }
    return hours + ' hours ago';
  }
  if (minutes > 0) {
    if (minutes == 1) {
      return '1 minute ago';
    }
    return minutes + ' minutes ago';
  }
  if (seconds > 0) {
    if (seconds === 1) {
      return '1 second ago';
    }
    return seconds + ' seconds ago';
  }
  return 'Just now';
}
